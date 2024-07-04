package views

import (
	"html/template"
)

func LoginPage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/loginPage.html"))
	return temp
}

func UserSignupPage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/registerPage.html"))
	return temp
}

func UserPage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/userPage.html"))
	return temp
}
