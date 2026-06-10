package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userId string) (string, error){

	claims := jwt.MapClaims{
		"sub": userId,
		"exp": time.Now().Add(24* time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	tokenString, err := token.SignedString(
		[]byte(os.Getenv("JWT_SECRET")),
	)
	if err != nil{
		return "", err
	}

	return tokenString, nil
}