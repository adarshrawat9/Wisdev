package server

import (
	"Wisdev/internal/database"
	"Wisdev/internal/handler"
	"Wisdev/internal/integrations/github"
	"Wisdev/internal/repositories"
	"Wisdev/internal/routes"
	"Wisdev/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func New()*gin.Engine{
	r := gin.Default()


	// create dependency
	userRepo := repositories.NewUserRepository(database.DB)
	userService := services.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// github integration  dependency
	githubClient := github.NewClient()
	githubServices := github.NewService(githubClient)
	githubHandler := handler.NewGithubHandler(githubServices)


	// register routes
	routes.RegisterUserRoutes(r, userHandler)
	routes.RegisterGithubRoutes(r, githubHandler)

	r.GET("/health", func (c *gin.Context)  {
		c.JSON(http.StatusOK, gin.H{"staus":"ok"})
	})



	

	return r
}