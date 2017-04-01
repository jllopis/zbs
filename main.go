// Copyright 2017 Joan Llopis. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/jllopis/getconf"
	"github.com/jllopis/zbs/log"
	"github.com/jllopis/zbs/metrics"
	"github.com/jllopis/zbs/server"
)

// Config proporciona la configuración del servicio para ser utilizado por getconf
type Config struct {
	GwAddr      string `getconf:"gwaddr, default 127.0.0.1, info http rest gateway address"`
	Port        string `getconf:"port, default 8000, info default port to listen on"`
	StoreDriver string `getconf:"store-driver, default boltdb, info storage backend"`
	StoreHost   string `getconf:"store-host, default db.acb.info, info store server address"`
	StorePort   string `getconf:"store-port, default 5432, info store server port"`
	StoreName   string `getconf:"store-name, default m3skeldb, info store database name"`
	StoreUser   string `getconf:"store-user, default m3skeladm, info store user to connect to db"`
	StorePass   string `getconf:"store-pass, default 00000000, info store user password"`
}

var (
	// BuildDate holds the date the binary was built. It is valued at compile time
	BuildDate string
	// Version holds the version number of the build. It is valued at compile time
	Version string
	// Revision holds the git revision of the binary. It is valued at compile time
	Revision string
	config   *getconf.GetConf
)

func main() {
	defer func() {
		if e := recover(); e != nil {
			stack := make([]byte, 1<<16)
			stackSize := runtime.Stack(stack, true)
			log.Info(string(stack[:stackSize]))
		}
	}()

	// load config (jllopis/getconf)
	config = getconf.New("ZBS", &Config{})

	// capture signals
	setupSignals()

	log.Info("ZBS service started", "port", config.Get("port"))
	log.Info("version", Version, "revision", Revision, "build-date", BuildDate)
	log.Info("GetConf", "version", getconf.Version())
	log.Info("Go", "version", runtime.Version())

	// config metrics
	metrics.StartMetrics()

	mssrv := server.New(config)
	mssrv.Register()

	log.Err(mssrv.Serve().Error())
}

// setupSignals configura la captura de señales de sistema y actúa basándose en ellas
func setupSignals() {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		for sig := range sc {
			log.Info("Capturada señal", sig)
			os.Exit(1)
		}
	}()
}

func printLicense() {
	fmt.Println(`ZBS  Copyright (C) 2015-2017  Joan Llopis
    This program comes with ABSOLUTELY NO WARRANTY.
    This is free software, and you are welcome to redistribute it
    under certain conditions; type 'show license' for details.`)
}
