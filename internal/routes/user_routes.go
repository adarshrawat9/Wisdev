package routes

import (
	"Wisdev/internal/handler"

	"github.com/gin-gonic/gin"
)


func RegisterUserRoutes(
	router *gin.Engine,
	userHandler *handler.UserHandler,
){
	auth := router.Group("/auth")
	{
		auth.POST("/register", userHandler.SignIn)
	}
}