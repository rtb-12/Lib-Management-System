package controller

import (
	"net/http"

	"github.com/rtb-12/Lib-Management-System/pkg/views"
)

func User(writer http.ResponseWriter, request *http.Request) {
	// id := request.URL.Query().Get("id")
	// if id == "" {
	// 	http.Error(writer, "Missing id query parameter", http.StatusBadRequest)
	// 	return
	// }
	t := views.UserPage()
	t.Execute(writer, nil)
}
