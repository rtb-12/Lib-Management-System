package views

import (
	"html/template"
)

func AdminLoginPage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/adminLogin.html"))
	return temp
}
