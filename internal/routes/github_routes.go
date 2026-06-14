package routes

import "github.com/gin-gonic/gin"


func RegisterGithubRoutes(
	routes *gin.Engine,
	githubHandler gin.HandlerFunc,
){
	routes.GET("/github/:username", githubHandler)
}

