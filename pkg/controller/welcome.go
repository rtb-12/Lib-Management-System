package controller

import (
	"net/http"

	"github.com/rtb-12/Lib-Management-System/pkg/views"
)

func Welcome(writer http.ResponseWriter, request *http.Request) {
	t := views.StartPage()
	t.Execute(writer, nil)
}
