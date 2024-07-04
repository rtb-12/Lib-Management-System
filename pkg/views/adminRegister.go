package views

import (
	"html/template"
)

func AdminRegisterPage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/adminRegister.html"))
	return temp
}
