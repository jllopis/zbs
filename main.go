package main

import (
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"bitbucket.org/jllopis/getconf"
	"github.com/boltdb/bolt"
	"github.com/golang/glog"
	"github.com/jllopis/aloja"
	"github.com/jllopis/aloja/mw"
)

// Config proporciona la configuración del servicio para ser utilizado por getconf
type Config struct {
	Port      string `getconf:"etcd app/zbs/conf/port, env ZBS_PORT, flag port"`
	Verbose   bool   `getconf:"etcd app/zbs/conf/verbose, env ZBS_VERBOSE, flag verbose"`
	StoreName string `getconf:"etcd app/zbs/conf/storename, env ZBS_STORE_NAME, flag storename"`
}

var (
	BUILD_DATE string
	VERSION    string
	REVISION   string
	config     *getconf.GetConf
	db         *bolt.DB
	verbose    bool
)

func init() {
	//etcdURI := os.Getenv("ZBS_ETCD")
	//config = getconf.New(&Config{}, "ZBS", true, etcdURI)
	config = getconf.New(&Config{}, "ZBS", false, "")
	config.Parse()
	var err error
	db, err = bolt.Open(config.GetString("StoreName"), 0600, nil)
	if err != nil {
		glog.Fatal(err)
	}
}

func main() {
	setupSignals()
	port := config.GetString("Port")
	if port == "" {
		glog.Infof("[zbs] can't get Port value from config")
		glog.Infof("[zbs] using default port 8000")
		port = "8000"
	}
	glog.Infof("ZBS service started on port: %s", port)
	glog.Infof("Version: %s", VERSION)
	glog.Infof("Revision: %s ( %s )", REVISION, BUILD_DATE)
	glog.Infof("GetConf Version: %s", getconf.Version())
	glog.Infof("Go Version: %s", runtime.Version())

	server := aloja.New().Port(port)
	// Use CORS Handler in every request and log every request
	server.AddGlobal(mw.CorsHandler, mw.LogHandler)

	// serve the angularjs app from /app subdirectory
	//app := server.NewSubrouter("/app")
	//app.ServeStatic("/", "/resources/")

	// run the server
	server.Run()
}

// setupSignals configura la captura de señales de sistema y actúa basándose en ellas
func setupSignals() {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		for sig := range sc {
			glog.Infof("[zbs.signal_notify] captured signal %v, stopping and exiting..", sig)
			os.Exit(1)
		}
	}()
}
