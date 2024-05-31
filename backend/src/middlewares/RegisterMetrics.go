package middlewares

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

const groupIndex = 3

var (
	apiRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "api_request_counter",
			Help: "API request times",
		},
		[]string{
			"group_name",
			"method",
		},
	)
	apiDurations = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "api_duration_seconds",
			Help:    "API request duration in seconds.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{
			"group_name",
			"method",
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

		path := c.FullPath()
		method := c.Request.Method

		if path != "/metrics" && !strings.HasPrefix(path, "/swagger") {
			groupName := strings.Split(path, "/")[groupIndex]
			apiRequests.WithLabelValues("/"+groupName, method).Inc()
			apiDurations.WithLabelValues(groupName, method).Observe(duration)
		}
	}
}
