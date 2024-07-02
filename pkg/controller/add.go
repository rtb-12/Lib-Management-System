package controller

import (
	"encoding/json"
	"fmt"
	"net/http"


	"github.com/ashpect/go-mvc/pkg/models"
	"github.com/ashpect/go-mvc/pkg/types"
)

func Add(writer http.ResponseWriter, request *http.Request) {
	var body types.Book
	err := json.NewDecoder(request.Body).Decode(&body)
	if err != nil {
		fmt.Print("There was an error decoding the request body into the struct")
	}

	fmt.Printf("Adding the book %s to the database\n", body.Name)
	models.AddBook(body.Name)
}
