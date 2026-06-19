package routes

import (
	"Wisdev/internal/handler"

	"github.com/gin-gonic/gin"
)


func RegisterGithubRoutes(
	router *gin.Engine,
	githubHandler *handler.GithubHandler,
){
	router.GET("/github/:username", githubHandler.GetUser)
	router.GET("/github/:username/repo", githubHandler.GetUserRepositories)
	router.GET("github/:username/analytics", githubHandler.GetUserAnalytics)
	router.GET("/github/:username/contributions", githubHandler.GetUserContributions)
}

