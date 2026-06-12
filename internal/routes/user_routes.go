package routes

import (
	"Wisdev/internal/handler"
	"Wisdev/internal/middleware"

	"github.com/gin-gonic/gin"
)


func RegisterUserRoutes(
	router *gin.Engine,
	userHandler *handler.UserHandler,
){
	auth := router.Group("/auth")
	{
		auth.POST("/register", userHandler.SignUp)
		auth.POST("/login", userHandler.Login)

		// protected routes needs to validate jwt
		protected := auth.Group("/")
		protected.Use(middleware.AuthMiddleware)
		protected.GET("/me", userHandler.Me)
	}

	users := router.Group("/users")
	{
		users.GET("/:username", userHandler.GetPublicProfile)
	}
}
