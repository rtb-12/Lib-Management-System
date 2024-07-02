package views

import (
	"html/template"
)

func UserPage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/userPage.html"))
	return temp
}
