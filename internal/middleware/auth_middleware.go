package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(c *gin.Context){
	authHeader := c.GetHeader("Authorization")

	parts := strings.SplitN(authHeader, " ", 2)

	// failed cause header doesn't split well
	if authHeader == "" || len(parts) != 2 || parts[0] != "Bearer"{
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error":"unauthorized",
		})
		return
	}

	//failed while verifying the token
	token, err := jwt.Parse(parts[1], func(token *jwt.Token) (interface{}, error){
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok{
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	// jwt parsing error 
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error":"unauthorized",
		})
		return

	}

	// token was invalid
	if !token.Valid{
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error":"unauthorized",
		})
		 return

	}

	claims, ok := token.Claims.(jwt.MapClaims)
	userId, uidok := claims["sub"].(string)


	if !ok || !uidok{
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})

	}

	c.Set("userId", userId)
	c.Next()
}