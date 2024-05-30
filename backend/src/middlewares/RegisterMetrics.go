package middlewares

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

const groupIndex = 2

var (
	apiRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "api_request_counter",
			Help: "API request times",
		},
		[]string{
			"user",
			"auth",
			"code",
			"room",
			"admin",
		},
	)
	apiDurations = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "api_duration_seconds",
			Help:    "API request duration in seconds.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{
			"user",
			"auth",
			"code",
			"room",
			"admin",
		},
	)
)

func init() {
	prometheus.MustRegister(apiRequests, apiDurations)
}

func RegisterMetricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		duration := time.Since(start).Seconds()

		group := strings.Split(c.FullPath(), "/")[groupIndex]
		method := c.Request.Method

		apiRequests.WithLabelValues(group, method).Inc()
		apiDurations.WithLabelValues(group, method).Observe(duration)
	}
}
