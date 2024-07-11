package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rtb-12/Lib-Management-System/pkg/models"
	"github.com/rtb-12/Lib-Management-System/pkg/types"
	"github.com/rtb-12/Lib-Management-System/pkg/views"
)

func Login(writer http.ResponseWriter, request *http.Request) {
	var userLogin types.UserLogin
	err := json.NewDecoder(request.Body).Decode(&userLogin)
	if err != nil {
		http.Error(writer, "Error decoding JSON body", http.StatusBadRequest)
		return
	}

	var userId int
	userId, err = models.CheckUser(userLogin)
	if err != nil {
		http.Error(writer, fmt.Sprintf("Error finding user: %s", err), http.StatusInternalServerError)
		return
	}

	// Generate JWT token
	secretKey := os.Getenv("JWT_SECRET")
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": strconv.Itoa(userId),
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})
	token, err := claims.SignedString([]byte(secretKey))
	if err != nil {
		fmt.Println("Error generating token:", err)
		http.Error(writer, fmt.Sprintf("Error generating token: %s", err), http.StatusInternalServerError)
		return
	}

	// Set JWT token in cookie
	expiration := time.Now().Add(24 * time.Hour)
	cookie := http.Cookie{
		Name:    "jwt",
		Value:   token,
		Expires: expiration,
		// HttpOnly: true,
		// Secure:   true,
	}
	http.SetCookie(writer, &cookie)

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	response := fmt.Sprintf(`{"message": "User logged in with jwt token", "token": "%s", "userID": "%d"}`, token, userId)
	writer.Write([]byte(response))
}

func LoginPage(writer http.ResponseWriter, request *http.Request) {
	t := views.LoginPage()
	t.Execute(writer, nil)
}
