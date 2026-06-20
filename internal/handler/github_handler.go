package handler

import (
	"Wisdev/internal/integrations/github"
	"net/http"

	"github.com/gin-gonic/gin"
)


type GithubHandler struct{
	service *github.Service
}

func NewGithubHandler(service *github.Service) *GithubHandler{
	return &GithubHandler{
		service: service,
	}
}

func (h *GithubHandler) GetUser(c *gin.Context){
	username := c.Param("username")

	user, err := h.service.GetUser(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)

}

func (h *GithubHandler) GetUserRepositories(c *gin.Context){

	username := c.Param("username")

	repositories, err := h.service.GetUserRepositories(username)
	if err != nil{
		c.JSON(http.StatusInternalServerError, 
		gin.H{
			"error" : err.Error(),
		})
		return
	}

	c.JSON(
		http.StatusOK,
		repositories,
	)


}

func (h *GithubHandler) GetUserAnalytics(
	c *gin.Context,
) {

	username := c.Param("username")

	analytics, err := h.service.GetUserAnalytics(
		username,
	)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		analytics,
	)
}


func (h *GithubHandler) GetUserContributions(c *gin.Context){

	username := c.Param("username")

	contributions, err := h.service.GetUserContributions(username)
	if err != nil{
		c.JSON(http.StatusInternalServerError,
		gin.H{
			"error": err.Error(),
		})
		return 
	}
	c.JSON(
		http.StatusOK,
		contributions,
	)
}

func (h *GithubHandler) GetUserGrowthAnalytics(
	c *gin.Context,
) {

	username := c.Param("username")

	growth, err := h.service.GetUserGrowthAnalytics(
		username,
	)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		growth,
	)
}