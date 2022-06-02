package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
)

type Vecs struct {
	Seconds  *prometheus.HistogramVec
	Requests *prometheus.CounterVec
}

func RegisterMetrics(namespace string) *Vecs {
	metricSeconds := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: namespace,
		Subsystem: "requests",
		Name:      "duration_ms",
		Help:      "server requests duration(ms).",
		Buckets:   []float64{5, 10, 25, 50, 100, 250, 500, 1000},
	}, []string{"kind", "operation"})

	metricRequests := prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: "requests",
		Name:      "code_total",
		Help:      "The total number of processed requests",
	}, []string{"kind", "operation", "code", "reason"})

	prometheus.MustRegister(
		metricSeconds,
		metricRequests,
		collectors.NewGoCollector(collectors.WithGoCollections(collectors.GoRuntimeMetricsCollection)),
	)

	return &Vecs{
		Seconds:  metricSeconds,
		Requests: metricRequests,
	}
}
