package controller

import (
	"net/http"

	"github.com/rtb-12/Lib-Management-System/pkg/views"
)

func Admin(writer http.ResponseWriter, request *http.Request) {
	t := views.AdminPage()
	t.Execute(writer, nil)
}
