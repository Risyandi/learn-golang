package middleware

import (
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/rs/zerolog"
	"github.com/shirou/gopsutil/v3/process"

	"gitlab.com/sugaanaluam/gofiber-boilerplate/config"
	"gitlab.com/sugaanaluam/gofiber-boilerplate/logger"
)

type (
	PrometheusMiddleware interface {
		RecordMetrics()
		MonitoringMiddleware() fiber.Handler
	}

	prometheusMiddleware struct {
		conf *config.Config
		log  zerolog.Logger
	}
)

// Define Prometheus metrics
var (
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "endpoint", "status"},
	)
	httpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "endpoint"},
	)
	httpRequestErrors = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_request_errors_total",
			Help: "Total number of HTTP request errors",
		},
		[]string{"method", "endpoint"},
	)
	cpuUsage = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "app_cpu_usage_seconds_total",
		Help: "Total CPU usage of the application",
	})
	cpuUsagePercent = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "app_cpu_usage_percent",
		Help: "CPU usage of the application in percent",
	})
	memoryUsage = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "app_memory_usage_bytes",
		Help: "Memory usage of the application",
	})
	memoryUsagePercent = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "app_memory_usage_percent",
		Help: "Memory usage of the application in percent",
	})
)

func NewPrometheusMiddleware() PrometheusMiddleware {
	return &prometheusMiddleware{
		conf: config.Get(),
		log:  logger.Get("prometheus-middleware"),
	}
}

func (m prometheusMiddleware) RecordMetrics() {
	prometheus.MustRegister(httpRequestsTotal, httpRequestDuration, httpRequestErrors, cpuUsage, cpuUsagePercent, memoryUsage, memoryUsagePercent)

	go func() {
		var cpuStats runtime.MemStats
		for {
			// Get process info
			proc, err := process.NewProcess(int32(os.Getpid())) // Get the current process
			if err != nil {
				m.log.Err(err).Msg("Failed Get Process ID")
			}

			// Collect memory usage
			runtime.ReadMemStats(&cpuStats)
			memoryUsage.Set(float64(cpuStats.Alloc))

			// Get Memory percentage
			memoryPercent, err := proc.MemoryPercent()
			if err != nil {
				m.log.Err(err).Msg("Failed fetching mem percent")
			} else {
				memoryUsagePercent.Set(float64(memoryPercent))
			}

			// Get CPU percentage
			cpuPercent, err := proc.CPUPercent()
			if err != nil {
				m.log.Err(err).Msg("Failed fetching CPU percent")
			} else {
				cpuUsagePercent.Set(cpuPercent)
			}

			// Update every 5 seconds
			time.Sleep(10 * time.Second)
		}
	}()
}

// Prometheus middleware
func (m prometheusMiddleware) MonitoringMiddleware() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		start := time.Now()

		// Process the request
		err := ctx.Next()

		// Track the metrics
		duration := time.Since(start).Seconds()
		httpRequestsTotal.WithLabelValues(ctx.Method(), ctx.Path(), http.StatusText(ctx.Response().StatusCode())).Inc()
		httpRequestDuration.WithLabelValues(ctx.Method(), ctx.Path()).Observe(duration)

		// Count errors if any
		if err != nil {
			httpRequestErrors.WithLabelValues(ctx.Method(), ctx.Path()).Inc()
		}

		return err
	}
}
