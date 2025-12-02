package controllers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joey17520/magic-stream-app/database"
)

// HealthCheck 返回应用健康状态
// @Summary 应用健康检查
// @Description 检查应用是否正常运行
// @Tags 健康检查
// @Produce json
// @Success 200 {object} map[string]interface{} "应用健康状态"
// @Router /health [get]
func HealthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":    "healthy",
			"timestamp": time.Now().Unix(),
			"service":   "magic-stream-api",
			"version":   "1.0.0",
		})
	}
}

// ReadyCheck 返回应用就绪状态
// @Summary 应用就绪检查
// @Description 检查应用是否就绪接收流量（检查数据库连接等）
// @Tags 健康检查
// @Produce json
// @Success 200 {object} map[string]interface{} "应用就绪状态"
// @Failure 503 {object} map[string]interface{} "应用未就绪"
// @Router /ready [get]
func ReadyCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 检查数据库连接
		if err := checkDatabaseConnection(); err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"status":    "not ready",
				"timestamp": time.Now().Unix(),
				"service":   "magic-stream-api",
				"error":     err.Error(),
				"checks": map[string]interface{}{
					"database": "unhealthy",
				},
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":    "ready",
			"timestamp": time.Now().Unix(),
			"service":   "magic-stream-api",
			"checks": map[string]interface{}{
				"database": "healthy",
			},
		})
	}
}

// LivenessCheck Kubernetes存活探针
// @Summary 应用存活检查
// @Description Kubernetes存活探针端点
// @Tags 健康检查
// @Produce json
// @Success 200 {object} map[string]interface{} "应用存活状态"
// @Router /live [get]
func LivenessCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":    "alive",
			"timestamp": time.Now().Unix(),
			"service":   "magic-stream-api",
		})
	}
}

// checkDatabaseConnection 检查数据库连接状态
func checkDatabaseConnection() error {
	// 尝试ping数据库
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 使用全局的Client变量
	if database.Client == nil {
		return errors.New("database client is nil")
	}

	// 执行ping操作检查连接
	err := database.Client.Ping(ctx, nil)
	if err != nil {
		return fmt.Errorf("database ping failed: %v", err)
	}

	return nil
}
