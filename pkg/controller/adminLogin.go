package controller

import (
	"net/http"

	"github.com/rtb-12/Lib-Management-System/pkg/views"
)

func AdminLogin(writer http.ResponseWriter, request *http.Request) {
	t := views.AdminLoginPage()
	t.Execute(writer, nil)
}
