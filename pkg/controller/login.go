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
	secret_key := os.Getenv("JWT_SECRET")
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": strconv.Itoa(int(userId)),
		"exp":    time.Now().Add(time.Hour * 24).Unix(), // Expires in 24 hours
	})
	token, err := claims.SignedString([]byte(secret_key))
	if err != nil {
		fmt.Println("Error generating token:", err)
		http.Error(writer, fmt.Sprintf("Error generating token : %s", err), http.StatusInternalServerError)
		return
	}

	// Set JWT token in cookie
	expiration := time.Now().Add(24 * time.Hour)
	cookie := http.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  expiration,
		HttpOnly: true,
		Secure:   true,
	}
	http.SetCookie(writer, &cookie)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	response := fmt.Sprintf(`{"message": "user logged in with jwt token", "token": "%s"}`, token)
	writer.Write([]byte(response))
}

func LoginPage(writer http.ResponseWriter, request *http.Request) {
	t := views.LoginPage()
	t.Execute(writer, nil)
}
