package controller

import (
	"net/http"

	"github.com/ashpect/go-mvc/pkg/models"
	"github.com/ashpect/go-mvc/pkg/views"
)

func List(writer http.ResponseWriter, request *http.Request) {
	// writer.WriteHeader(http.StatusOK)
	t := views.ListPage()
	writer.WriteHeader(http.StatusOK)
	booksList := models.FetchBooks()
	t.Execute(writer, booksList)
}
