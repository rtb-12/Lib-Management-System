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

func AdminLogin(writer http.ResponseWriter, request *http.Request) {

	var adminLogin types.AdminLogin
	err := json.NewDecoder(request.Body).Decode(&adminLogin)
	if err != nil {
		http.Error(writer, "Error decoding JSON body", http.StatusBadRequest)
		return
	}
	var adminID int
	adminID, err = models.CheckAdmin(adminLogin)
	if err != nil {
		http.Error(writer, fmt.Sprintf("Error finding admin: %s", err), http.StatusInternalServerError)
		return
	}

	// Generate JWT token
	secret_key := os.Getenv("JWT_SECRET")
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"adminID": strconv.Itoa(int(adminID)),
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Expires in 24 hours
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
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	response := fmt.Sprintf(`{"message": "amdin logged in with jwt token", "token": "%s"}`, token)
	writer.Write([]byte(response))
}

func AdminLoginPage(writer http.ResponseWriter, request *http.Request) {
	t := views.AdminLoginPage()
	t.Execute(writer, nil)
}
