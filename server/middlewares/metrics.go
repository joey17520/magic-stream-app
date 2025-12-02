package middlewares

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	// HTTP请求总数
	httpRequestsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "path", "status"},
	)

	// HTTP请求持续时间
	httpRequestDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "HTTP request duration in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "path"},
	)

	// HTTP请求大小
	httpRequestSize = promauto.NewSummaryVec(
		prometheus.SummaryOpts{
			Name: "http_request_size_bytes",
			Help: "HTTP request size in bytes",
		},
		[]string{"method", "path"},
	)

	// HTTP响应大小
	httpResponseSize = promauto.NewSummaryVec(
		prometheus.SummaryOpts{
			Name: "http_response_size_bytes",
			Help: "HTTP response size in bytes",
		},
		[]string{"method", "path"},
	)

	// 活跃请求数
	httpRequestsInFlight = promauto.NewGauge(
		prometheus.GaugeOpts{
			Name: "http_requests_in_flight",
			Help: "Current number of HTTP requests being served",
		},
	)

	// 数据库操作指标
	dbOperationsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "db_operations_total",
			Help: "Total number of database operations",
		},
		[]string{"operation", "collection"},
	)

	dbOperationDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "db_operation_duration_seconds",
			Help:    "Database operation duration in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"operation", "collection"},
	)

	// 业务特定指标
	moviesViewedTotal = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "movies_viewed_total",
			Help: "Total number of movies viewed",
		},
	)

	usersRegisteredTotal = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "users_registered_total",
			Help: "Total number of users registered",
		},
	)

	recommendationsGeneratedTotal = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "recommendations_generated_total",
			Help: "Total number of recommendations generated",
		},
	)
)

// MetricsMiddleware 收集HTTP请求指标
func MetricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 跳过健康检查端点的指标收集
		path := c.Request.URL.Path
		if path == "/metrics" || path == "/health" || path == "/ready" || path == "/live" {
			c.Next()
			return
		}

		// 记录请求开始时间和活跃请求数
		start := time.Now()
		httpRequestsInFlight.Inc()

		// 记录请求大小
		requestSize := float64(c.Request.ContentLength)
		if requestSize < 0 {
			requestSize = 0
		}
		httpRequestSize.WithLabelValues(c.Request.Method, path).Observe(requestSize)

		// 处理请求
		c.Next()

		// 记录响应大小
		responseSize := float64(c.Writer.Size())
		httpResponseSize.WithLabelValues(c.Request.Method, path).Observe(responseSize)

		// 记录请求持续时间
		duration := time.Since(start).Seconds()
		httpRequestDuration.WithLabelValues(c.Request.Method, path).Observe(duration)

		// 记录请求总数
		status := strconv.Itoa(c.Writer.Status())
		httpRequestsTotal.WithLabelValues(c.Request.Method, path, status).Inc()

		// 减少活跃请求数
		httpRequestsInFlight.Dec()
	}
}

// RecordDBOperation 记录数据库操作指标
func RecordDBOperation(operation, collection string, duration time.Duration) {
	dbOperationsTotal.WithLabelValues(operation, collection).Inc()
	dbOperationDuration.WithLabelValues(operation, collection).Observe(duration.Seconds())
}

// RecordMovieViewed 记录电影观看事件
func RecordMovieViewed() {
	moviesViewedTotal.Inc()
}

// RecordUserRegistered 记录用户注册事件
func RecordUserRegistered() {
	usersRegisteredTotal.Inc()
}

// RecordRecommendationGenerated 记录推荐生成事件
func RecordRecommendationGenerated() {
	recommendationsGeneratedTotal.Inc()
}

// GetMetricsHandler 返回Prometheus指标处理器
func GetMetricsHandler() gin.HandlerFunc {
	// 创建Prometheus HTTP处理器
	handler := promhttp.Handler()

	return func(c *gin.Context) {
		// 使用Prometheus处理器处理/metrics端点
		handler.ServeHTTP(c.Writer, c.Request)
	}
}
