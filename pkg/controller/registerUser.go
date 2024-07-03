package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rtb-12/Lib-Management-System/pkg/models"
	"github.com/rtb-12/Lib-Management-System/pkg/types"
)

func Register(writer http.ResponseWriter, request *http.Request) {
	// Step 1: Decode the JSON body into the User struct
	var user types.User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		http.Error(writer, "Error decoding JSON body", http.StatusBadRequest)
		return
	}

	// Step 2: Call the RegisterUser function with the parsed User struct
	err = models.RegisterUser(user)
	if err != nil {
		// Handle errors from RegisterUser, e.g., user already exists or database issues
		http.Error(writer, fmt.Sprintf("Error registering user: %s", err), http.StatusInternalServerError)
		return
	}

	// Step 3: Respond to the client
	writer.WriteHeader(http.StatusCreated)
	writer.Write([]byte("User registered successfully"))
}
