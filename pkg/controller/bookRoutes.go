package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rtb-12/Lib-Management-System/pkg/models"
	"github.com/rtb-12/Lib-Management-System/pkg/types"
)

func Add(writer http.ResponseWriter, request *http.Request) {
	var book types.BookInfo
	err := json.NewDecoder(request.Body).Decode(&book)
	if err != nil {
		fmt.Print("There was an error decoding the request body into the struct")
	}

	isBookAdded, err := models.AddBook(book)
	if err != nil {
		fmt.Printf("Error adding book: %s\n", err)
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Header().Set("Content-Type", "application/json")
		response := fmt.Sprintf(`{"message": "Error Adding the Book probably already exist", "BookName": "%s", "BookID": "%d", "Author":"%s"  }`, book.Title, book.ID, book.Author)
		writer.Write([]byte(response))
		return
	}
	if isBookAdded {
		writer.WriteHeader(http.StatusOK)
		writer.Header().Set("Content-Type", "application/json")
		response := fmt.Sprintf(`{"message": "Book Added", "BookName": "%s", "BookID": "%d", "Author":"%s"  }`, book.Title, book.ID, book.Author)
		writer.Write([]byte(response))
	} else {
		writer.WriteHeader(http.StatusBadRequest)
	}
}

func UpdateBook(writer http.ResponseWriter, request *http.Request) {
	var book types.BookInfo
	err := json.NewDecoder(request.Body).Decode(&book)
	if err != nil {
		fmt.Print("There was an error decoding the request body into the struct")
	}

	isBookUpdated, err := models.UpdateBook(book)
	if err != nil {
		fmt.Printf("Error adding book: %s\n", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	if isBookUpdated {
		writer.WriteHeader(http.StatusOK)
		writer.Header().Set("Content-Type", "application/json")
		response := fmt.Sprintf(`{"message": "Book Updated", "BookName": "%s", "BookID": "%d", "Author":"%s" , "Genre":"%s" , "Stock": "%d"  }`, book.Title, book.ID, book.Author, book.Genre, book.Stock)
		writer.Write([]byte(response))
	} else {
		writer.WriteHeader(http.StatusBadRequest)
	}
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	var book types.BookInfo
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		fmt.Print("There was an error decoding the request body into the struct")
	}

	isBookDeleted, err := models.DeleteBook(book)
	if err != nil {
		fmt.Printf("Error deleting book: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if isBookDeleted {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		response := fmt.Sprintf(`{"message": "Book Deleted", "BookName": "%s", "BookID": "%d", "Author":"%s" , "Genre":"%s" , "Stock": "%d"  }`, book.Title, book.ID, book.Author, book.Genre, book.Stock)
		w.Write([]byte(response))
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}
