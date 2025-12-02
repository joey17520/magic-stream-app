package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/joey17520/magic-stream-app/controllers"
	"github.com/joey17520/magic-stream-app/middlewares"
)

func SetupUnprotectedRoutes(router *gin.Engine) {
	// 健康检查端点
	router.GET("/health", controllers.HealthCheck())
	router.GET("/ready", controllers.ReadyCheck())
	router.GET("/live", controllers.LivenessCheck())

	// 指标端点
	router.GET("/metrics", middlewares.GetMetricsHandler())

	// 业务端点
	router.GET("/movies", controllers.GetMovies())
	router.POST("/register", controllers.RegisterUser())
	router.POST("/login", controllers.LoginUser())
	router.POST("/logout", controllers.LogoutHandler())
	router.GET("/genres", controllers.GetGenres())
	router.POST("/refresh", controllers.RefreshTokenHandler())
}
