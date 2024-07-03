package utils

import (
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Assuming 'secretKey' is a global variable or constant that holds your JWT secret key.
var secretKey = "your_secret_key"

// GenerateJWTToken takes a userID and returns a JWT token string and any error encountered.
func GenerateJWTToken(userID int) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": strconv.Itoa(userID),
		"exp": time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	})

	token, err := claims.SignedString([]byte(secretKey))
	if err != nil {
		fmt.Println("Error generating token:", err)
		return "", err
	}

	return token, nil
}
