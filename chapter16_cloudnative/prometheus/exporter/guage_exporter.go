package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

var (
	opsQueued = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "our_company",
		Subsystem: "blob_storage",
		Name:      "ops_queued",
		Help:      "Number of blob storage operations waiting to be processed.",
	})
)

func init() {
	prometheus.MustRegister(opsQueued)
}

func main() {
	go func() {
		// 初始值10
		opsQueued.Add(10)
		// 减1
		opsQueued.Dec()
		// 减1
		opsQueued.Dec()
	}()

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
