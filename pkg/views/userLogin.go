package views

import (
	"html/template"
)

func LoginPage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/loginPage.html"))
	return temp
}
