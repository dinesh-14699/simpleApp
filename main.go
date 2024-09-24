package main

import (
	"net/http"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var requestCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total number of HTTP requests.",
	},
)

func main() {
	prometheus.MustRegister(requestCounter)

	http.Handle("/metrics", promhttp.Handler())

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requestCounter.Inc()
		w.Write([]byte("Hello, Prometheus!"))
	})

	http.ListenAndServe(":8080", nil)
}
