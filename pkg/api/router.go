package api

import (
	"net/http"

	"github.com/gorilla/mux"
	controller "github.com/rtb-12/Lib-Management-System/pkg/controller"
	"github.com/rtb-12/Lib-Management-System/pkg/middlewares"
)

func Start() {
	r := mux.NewRouter()
	r.HandleFunc("/", controller.Welcome).Methods("GET")
	r.HandleFunc("/add", controller.Add).Methods("POST")
	r.HandleFunc("/list", controller.List).Methods("GET")
	r.HandleFunc("/admin/login", controller.AdminLogin).Methods("POST")
	r.HandleFunc("/admin/login", controller.AdminLoginPage).Methods("GET")
	r.HandleFunc("/admin/register", controller.RegisterAdminstrator).Methods("POST")
	r.HandleFunc("/admin/register", controller.AdminRegisterPage).Methods("GET")
	adminRoute := r.PathPrefix("/admin").Subrouter()
	adminRoute.Use(middlewares.AdminAuthMiddleware)
	adminRoute.HandleFunc("", controller.Admin).Methods("GET")
	r.HandleFunc("/login", controller.Login).Methods("GET")
	r.HandleFunc("/delete", controller.Delete).Methods("DELETE")
	r.HandleFunc("/user", controller.User).Methods("GET")
	r.HandleFunc("/register", controller.Register).Methods("POST")
	// r.HandleFunc("/update", controller.Update).Methods("PUT")
	// r.HandleFunc("/requestIssues", controller.requestIssues).Methods("POST")
	// r.HandleFunc("/issueBook", controller.issueBook).Methods("POST")
	// r.HandleFunc("/returnBook", controller.returnBook).Methods("POST")
	// r.HandleFunc("/listIssues", controller.listIssues).Methods("GET")
	http.ListenAndServe("localhost:8000", r)
}
