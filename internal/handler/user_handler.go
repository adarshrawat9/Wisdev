package handler

import (
	"Wisdev/internal/dto"
	"Wisdev/internal/services"
	"log"
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

func (u *UserHandler)Login(c *gin.Context){

	var req dto.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"invalid request",
		})
		return
	}

	token, err := u.service.Login(req)
	if err != nil{
		
		if err.Error() == "Invalid credentials"{
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid credentials",
			})
			return
		}

		log.Printf("login failed: %v", err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"error":"something unexpected occured",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}


func (u *UserHandler) Me(c *gin.Context) {
	
	userId := c.GetString("userId")

	user, err := u.service.GetById(userId)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":"something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, dto.UserResponse{
		ID:               user.ID,
        Username:         user.Username,
        Email:            user.Email,
        Bio:              user.Bio,
        GithubUsername:   user.GithubUsername,
        PortfolioWebsite: user.PortfolioWebsite,
        AvatarURL:        user.AvatarURL,
	})

}