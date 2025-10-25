package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/joey17520/magic-stream-app/controllers"
	"github.com/joey17520/magic-stream-app/middlewares"
)

func SetupProtectedRoutes(router *gin.Engine) {
	router.Use(middlewares.AuthMiddleware())

	router.GET("/movie/:imdb_id", controllers.GetMovie())
	router.POST("/movie", controllers.AddMovie())
	router.GET("/recommendedmovies", controllers.GetRecommendedMovies())
	router.PATCH("/updatereview/:imdb_id", controllers.AdminReviewUpdate())
}
