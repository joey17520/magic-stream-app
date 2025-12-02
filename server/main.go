package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joey17520/magic-stream-app/config"
	"github.com/joey17520/magic-stream-app/database"
	"github.com/joey17520/magic-stream-app/middlewares"
	"github.com/joey17520/magic-stream-app/routes"
	"github.com/joey17520/magic-stream-app/utils"
	"go.uber.org/zap"
)

func main() {
	// 初始化结构化日志
	if err := utils.InitLogger(); err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}
	defer utils.SyncLogger()

	logger := utils.GetLogger()
	logger.Info("Initializing MagicStream server")

	// 加载配置
	cfg := config.LoadConfig(logger)

	// 初始化数据库连接
	logger.Info("Initializing database connection")
	if err := database.InitDB(cfg); err != nil {
		logger.Fatal("Failed to initialize database", zap.Error(err))
	}
	defer database.CloseDB()

	// 设置用户集合到utils包
	userCollection := database.OpenCollection("users")
	if userCollection != nil {
		utils.SetUserCollection(userCollection)
		logger.Debug("User collection initialized in utils package")
	}

	router := gin.New()

	// CORS配置
	corsConfig := cors.Config{
		AllowOrigins:     cfg.AllowedOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		MaxAge:           12 * time.Hour,
		AllowCredentials: true, // 携带http-only cookie
	}

	// 中间件顺序很重要：CORS -> 结构化日志 -> 指标
	router.Use(cors.New(corsConfig))
	router.Use(StructuredLogger(logger))
	router.Use(middlewares.MetricsMiddleware())

	routes.SetupUnprotectedRoutes(router)
	routes.SetupProtectedRoutes(router)

	logger.Info("Starting MagicStream server",
		zap.String("port", cfg.ServerPort),
		zap.Strings("allowed_origins", cfg.AllowedOrigins),
	)

	logger.Info("Health check endpoints available",
		zap.String("health_endpoint", "GET /health"),
		zap.String("ready_endpoint", "GET /ready"),
		zap.String("live_endpoint", "GET /live"),
		zap.String("metrics_endpoint", "GET /metrics"),
	)

	if err := router.Run(":" + cfg.ServerPort); err != nil {
		logger.Fatal("Failed to start server", zap.Error(err))
	}
}

// StructuredLogger 替换gin的默认日志中间件，使用zap结构化日志
func StructuredLogger(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 跳过健康检查端点的详细日志
		path := c.Request.URL.Path
		if path == "/metrics" || path == "/health" || path == "/ready" || path == "/live" {
			c.Next()
			return
		}

		start := time.Now()
		c.Next()
		end := time.Now()
		latency := end.Sub(start)

		fields := []zap.Field{
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("ip", c.ClientIP()),
			zap.Duration("latency", latency),
			zap.String("user_agent", c.Request.UserAgent()),
		}

		if len(c.Errors) > 0 {
			for _, e := range c.Errors.Errors() {
				logger.Error(e, fields...)
			}
		} else {
			switch {
			case c.Writer.Status() >= 500:
				logger.Error("Server error", fields...)
			case c.Writer.Status() >= 400:
				logger.Warn("Client error", fields...)
			default:
				logger.Info("Request completed", fields...)
			}
		}
	}
}
