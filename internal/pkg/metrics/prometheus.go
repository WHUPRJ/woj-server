package metrics

import (
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/WHUPRJ/woj-server/internal/pkg/cast"
	"github.com/prometheus/client_golang/prometheus"
)

type Metrics struct {
	namespace string
	subsystem string

	counter *prometheus.CounterVec
	hist    *prometheus.HistogramVec

	logPaths []string
}

func (m *Metrics) Setup(namespace string, subsystem string) {
	m.namespace = namespace
	m.subsystem = subsystem

	m.counter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "requests_total",
			Help:      "Total number of requests",
		},
		[]string{"method", "url"},
	)

	m.hist = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "requests_details",
			Help:      "Details of each request",
			Buckets:   []float64{0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1, 2.5, 5, 10, 25, 50, 100, 250, 500, 1000},
		},
		[]string{"method", "url", "success", "http_code", "err_code"},
	)

	prometheus.MustRegister(m.counter, m.hist)
}

func (m *Metrics) Record(method, url string, success bool, httpCode int, errCode e.Status, elapsed float64) {
	m.counter.With(prometheus.Labels{
		"method": method,
		"url":    url,
	}).Inc()

	m.hist.With(prometheus.Labels{
		"method":    method,
		"url":       url,
		"success":   cast.ToString(success),
		"http_code": cast.ToString(httpCode),
		"err_code":  cast.ToString(errCode),
	}).Observe(elapsed)
}
