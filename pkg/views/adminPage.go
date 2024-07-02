package views

import (
	"html/template"
)

func AdminPage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/adminPage.html"))
	return temp
}
