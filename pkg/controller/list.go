package controller

import (
	"encoding/json"
	"net/http"

	"github.com/rtb-12/Lib-Management-System/pkg/models"
)

// func List(writer http.ResponseWriter, request *http.Request) {
// 	t := views.ListPage()
// 	writer.WriteHeader(http.StatusOK)
// 	booksList := models.FetchBooks()
// 	t.Execute(writer, booksList)
// }

func List(writer http.ResponseWriter, request *http.Request) {
	booksList := models.FetchBooks() // Assuming FetchBooks returns a slice of books

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(writer).Encode(booksList); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}
