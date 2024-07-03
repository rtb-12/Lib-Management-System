package controller

import (
	"net/http"

	"github.com/rtb-12/Lib-Management-System/pkg/models"
	"github.com/rtb-12/Lib-Management-System/pkg/views"
)

func List(writer http.ResponseWriter, request *http.Request) {
	t := views.ListPage()
	writer.WriteHeader(http.StatusOK)
	booksList := models.FetchBooks()
	t.Execute(writer, booksList)
}
