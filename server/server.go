// Copyright 2017 Joan Llopis. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package server

import (
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/jllopis/zbs/log"
	svc "github.com/jllopis/zbs/services"
	impl "github.com/jllopis/zbs/services/implementation"
	"github.com/jllopis/zbs/store"
	"github.com/jllopis/zbs/store/mem"
	"github.com/prometheus/client_golang/prometheus"
	"golang.org/x/net/context"

	"github.com/grpc-ecosystem/go-grpc-prometheus"
	gw "github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/jllopis/getconf"
	"google.golang.org/grpc"

	"github.com/cockroachdb/cmux"
	"github.com/codahale/metrics"
)

type MicroServer struct {
	cmuxServer   cmux.CMux
	grpcListener net.Listener
	GrpcServer   *grpc.Server
	GrpcGwMux    *gw.ServeMux
	httpListener net.Listener
	httpServer   *http.Server

	config *getconf.GetConf

	defaultStore store.Storer
}

// New crea un nuevo MicroServer a partir de las opciones de configuración proporcionadas.
// Devuelve una referencia a MicroServer. La función siempre devuelve una estructura.
// Si se produce un error en su creación, se devuelve &MicroServer{}
func New(c *getconf.GetConf) *MicroServer {
	fmt.Printf("\n%v\n", c)
	// TODO (jllopis): permitir a través de una variable especificar si ha de conectarse con
	// un Store o no.
	// Conexión con el store
	// ds, err := gopg.New()
	ds, err := mem.New()
	if err != nil {
		log.Err("cannot create default store", "error", err)
	}
	// Dialing store
	if err := ds.Dial(store.Options{
		"host":   c.GetString("store-host"),
		"dbname": c.GetString("store-name"),
	}); err != nil {
		log.Err("cannot create default store", err.Error())
	}

	// Crear el MicroServer
	ms := &MicroServer{
		config:       c,
		defaultStore: ds,
	}

	// Create the main listener.
	l, err := net.Listen("tcp", ":"+ms.config.GetString("port"))
	if err != nil {
		log.Err(err.Error())
		os.Exit(1)
	}

	// Crear el muxer cmux
	ms.cmuxServer = cmux.New(l)

	// Match connections in order:
	// First grpc, and otherwise HTTP.
	ms.grpcListener = ms.cmuxServer.Match(cmux.HTTP2HeaderField("content-type", "application/grpc"))
	// Any significa que no hay coincidencia previa
	// En nuestro caso, no es grpc así que debe ser http.
	ms.httpListener = ms.cmuxServer.Match(cmux.Any())

	// Create your protocol servers.
	ms.GrpcServer = grpc.NewServer(
		grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
		grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor),
	)
	grpc_prometheus.EnableHandlingTimeHistogram()
	grpc_prometheus.Register(ms.GrpcServer)

	ms.GrpcGwMux = gw.NewServeMux()

	httpmux := http.NewServeMux()
	httpmux.Handle("/", logh(ms.GrpcGwMux))

	httpmux.Handle("/metrics", prometheus.Handler())

	ms.httpServer = &http.Server{
		Handler: httpmux,
		// ErrorLog: logger, // do not log user error
	}

	return ms
}

// Serve inica los servicios grpc y http para escuchar peticiones
func (ms *MicroServer) Serve() error {
	// Use the muxed listeners for your servers.
	go ms.GrpcServer.Serve(ms.grpcListener)
	go ms.httpServer.Serve(ms.httpListener)
	// Start serving!
	return ms.cmuxServer.Serve()

}

func (ms *MicroServer) Register() []error {

	errors := []error{}

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	// Register grpc endpoint on gRPC server
	svc.RegisterTimeServiceServer(ms.GrpcServer, &impl.ServerTime{})
	// Register REST endpoint on grpc-gw server
	if err := svc.RegisterTimeServiceHandlerFromEndpoint(context.Background(), ms.GrpcGwMux, ms.config.GetString("gwaddr")+":"+ms.config.GetString("port"), opts); err != nil {
		errors = append(errors, err)
	}

	// ZbsService
	svc.RegisterZbsServiceServer(ms.GrpcServer, &impl.ServiceZbs{Store: ms.defaultStore})
	if err := svc.RegisterZbsServiceHandlerFromEndpoint(context.Background(), ms.GrpcGwMux, ms.config.GetString("gwaddr")+":"+ms.config.GetString("port"), opts); err != nil {
		errors = append(errors, err)
	}

	// Devolver los posibles errores producidos en la inicialización de los servicios
	return errors
}

func logh(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info(fmt.Sprintf("%s %s %s", r.RemoteAddr, r.Method, r.URL))
		// Registrar llamada REST
		metrics.Counter("rest.requests").Add()
		handler.ServeHTTP(w, r)
	})
}
