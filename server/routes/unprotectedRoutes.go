package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/joey17520/magic-stream-app/controllers"
)

func SetupUnprotectedRoutes(router *gin.Engine) {
	router.GET("/movies", controllers.GetMovies())
	router.POST("/register", controllers.RegisterUser())
	router.POST("/login", controllers.LoginUser())
	router.POST("/logout", controllers.LogoutHandler())
	router.GET("/genres", controllers.GetGenres())
	router.POST("/refresh", controllers.RefreshTokenHandler())

}
