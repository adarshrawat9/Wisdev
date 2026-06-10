package handler

import (
	"Wisdev/internal/dto"
	"Wisdev/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)


type UserHandler struct{
	service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler{
	return &UserHandler{
		service: service,
	}
}

func (u *UserHandler)SignIn(c *gin.Context){
	
	var req dto.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"invalid request",
		})
		return
	}

	_, err := u.service.Register(req)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "user registered successfully",
	})

}