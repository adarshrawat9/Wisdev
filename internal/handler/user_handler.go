package handler

import (
	"Wisdev/internal/dto"
	"Wisdev/internal/services"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)


type UserHandler struct{
	service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler{
	return &UserHandler{
		service: service,
	}
}

func (u *UserHandler)SignUp(c *gin.Context){
	
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

func (u *UserHandler) UpdateUserDetails(c *gin.Context){

	userId := c.GetString("userId")

	var req dto.UpdateUserProfile

	if err := c.ShouldBindJSON(&req); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid values passed",
		})
		return
	}

	user, err := u.service.UpdateUserDetails(userId, req)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	userResponse := dto.UserResponse{
	    	ID: user.ID,
	    	Username: user.Username,
	    	Email: user.Email,
	    	Bio: user.Bio,
	    	GithubUsername: user.GithubUsername,
	    	PortfolioWebsite: user.PortfolioWebsite,
	    	AvatarURL: user.AvatarURL,
	    }

	c.JSON(http.StatusOK, userResponse)


}

func (u *UserHandler) GetPublicProfile(c *gin.Context){

	username := c.Param("username")

	user, err := u.service.GetPublicProfile(username)
	if err != nil{
		if errors.Is(err, pgx.ErrNoRows) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "user not found",
		})
		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{
		"error": "internal server error",
	})
	return
	}

	response := dto.PublicProfileResponse{
		Username: user.Username,
		Bio : user.Bio,
		GithubUsername: user.GithubUsername,
		PortfolioWebsite: user.PortfolioWebsite,
		AvatarURL: user.AvatarURL,
	}

	c.JSON(http.StatusOK, response)

}