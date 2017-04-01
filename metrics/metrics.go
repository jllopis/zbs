// Copyright 2017 Joan Llopis. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package metrics

import (
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/jllopis/zbs/log"

	"github.com/codahale/metrics"
	_ "github.com/codahale/metrics/runtime"
)

type metricSerie struct {
	Name    string        `json:"name"`
	Columns []string      `json:"columns"`
	Points  []metricPoint `json:"points"`
}

type metricPoint interface{}

var (
	statusSerie         metricSerie
	restRequestsSerie   metricSerie
	statsIntervalTicker = time.Duration(30)
	serverIP            string
)

func init() {
	if val := os.Getenv("ZBS_INTERVAL_TICKER"); val != "" {
		if t, err := time.ParseDuration(val); err == nil {
			statsIntervalTicker = t
		}
	}
	restRequestsSerie = metricSerie{
		"rest.requests",
		[]string{
			"host",
			"value",
		},
		make([]metricPoint, 2),
	}
	statusSerie = metricSerie{
		"Memory",
		[]string{
			"host",
			"MemNumGC",
			"MemPauseTotalNS",
			"MemHeapObjects",
			"FileDescriptorsMax",
			"FileDescriptorsUsed",
			"GoroutinesNum",
			"MemLastGC",
			"Mem.Alloc",
		},
		make([]metricPoint, 9),
	}
}

func StartMetrics() {
	interfaces, _ := net.Interfaces()
	for _, inter := range interfaces {
		if addresses, err := inter.Addrs(); err == nil {
			for _, addr := range addresses {
				if addr.String() != "" && !strings.HasPrefix(addr.String(), "127.") {
					serverIP = addr.String()
					break
				}
			}
		} else {
			serverIP = "UNKNOWN"
		}
	}
	// metrics server
	go http.ListenAndServe(":8123", nil)
	log.Info("Started monitoring endpoint", "host", "localhost", "port", "8123")

	go func() {
		c := time.Tick(statsIntervalTicker * time.Second)
		for range c {
			var counters map[string]uint64
			var gauges map[string]int64
			counters, gauges = metrics.Snapshot()
			restRequestsSerie.Points = []metricPoint{
				serverIP,
				counters["rest.requests"],
			}
			statusSerie.Points = []metricPoint{
				serverIP,
				counters["Mem.NumGC"],
				counters["Mem.PauseTotalNs"],
				gauges["Mem.HeapObjects"],
				gauges["FileDescriptors.Max"],
				gauges["FileDescriptors.Used"],
				gauges["Goroutines.Num"],
				gauges["Mem.LastGC"],
				gauges["Mem.Alloc"],
			}
			//			ss, err := json.Marshal(restRequestsSerie)
			//			if err != nil {
			//				log.Printf("metrics: %v", err)
			//			}
			//			ms, err := json.Marshal(statusSerie)
			//			if err != nil {
			//				log.Printf("metrics: %v", err)
			//			}
			//log.Info("rest.requests", counters["rest.requests"], "mem.alloc", gauges["Mem.Alloc"])
		}
	}()
}
