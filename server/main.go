package main

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joey17520/magic-stream-app/routes"
)

func main() {
	router := gin.New()

	// 使用http-only Cookie必须指定信任源
	allowedOrigins := os.Getenv("ALLOWED_ORIGINS")
	var origins []string
	if allowedOrigins != "" {
		origins = strings.Split(allowedOrigins, ",")
		for i := range origins {
			origins[i] = strings.TrimSpace(origins[i])
		}
	} else {
		origins = []string{"http://localhost:5173"}
	}

	config := cors.Config{
		AllowOrigins:     origins,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		MaxAge:           12 * time.Hour,
		AllowCredentials: true, // 携带http-only cookie
	}

	router.Use(cors.New(config))
	router.Use(gin.Logger())

	routes.SetupUnprotectedRoutes(router)
	routes.SetupProtectedRoutes(router)

	if err := router.Run(":8088"); err != nil {
		log.Fatal("Failed to start server ", err)
	}
}
