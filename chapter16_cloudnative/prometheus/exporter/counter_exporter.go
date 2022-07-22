package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"time"
)

var (
	demoCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "AAA",
		Subsystem: "BBB",
		Name:      "CCC",
		Help:      "this is a counter demo",
	})
)

func init() {
	prometheus.Register(demoCounter)
}
func main() {
	go func() {
		for {
			demoCounter.Inc()
			time.Sleep(time.Second * 10)
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
