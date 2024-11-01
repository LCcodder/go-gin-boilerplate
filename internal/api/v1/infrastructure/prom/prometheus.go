package prom

import "github.com/prometheus/client_golang/prometheus"

var (
	UserCreatedCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "user_created_count",
			Help: "Number of users created.",
		},
		[]string{"method"},
	)
)

func RegisterPrometheusMetrics() {
	prometheus.MustRegister(UserCreatedCounter)
}