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

	r.HandleFunc("/list", controller.List).Methods("GET")
	r.HandleFunc("/admin/login", controller.AdminLogin).Methods("POST")
	r.HandleFunc("/admin/login", controller.AdminLoginPage).Methods("GET")
	r.HandleFunc("/admin/register", controller.RegisterAdminstrator).Methods("POST")
	r.HandleFunc("/admin/register", controller.AdminRegisterPage).Methods("GET")
	adminRoute := r.PathPrefix("/admin").Subrouter()
	adminRoute.Use(middlewares.AdminAuthMiddleware)
	adminRoute.HandleFunc("", controller.Admin).Methods("GET")
	// adminRoute.HandleFunc("/add", controller.Add).Methods("POST")
	// adminRoute.HandleFunc("/updatebook", controller.UpdateBook).Methods("PUT")
	// adminRoute.HandleFunc("/delete", controller.Delete).Methods("DELETE")
	r.HandleFunc("/addbook", controller.Add).Methods("POST")
	r.HandleFunc("/updatebook", controller.UpdateBook).Methods("PUT")
	r.HandleFunc("/deletebook", controller.DeleteBook).Methods("DELETE")
	r.HandleFunc("/login", controller.LoginPage).Methods("GET")
	r.HandleFunc("/login", controller.Login).Methods("POST")

	userRoute := r.PathPrefix("/user").Subrouter()
	userRoute.Use(middlewares.LoginAuthMiddleware)
	userRoute.HandleFunc("", controller.User).Methods("GET")
	r.HandleFunc("/signup", controller.Register).Methods("POST")
	r.HandleFunc("/signup", controller.RegisterPage).Methods("GET")
	r.HandleFunc("/requestbookissue", controller.RequestBookIssue).Methods("POST")
	r.HandleFunc("/allocatebook", controller.IssueBook).Methods("POST")
	r.HandleFunc("/bookreturn", controller.ReturnBookRequest).Methods("POST")
	r.HandleFunc("/rejectbookissue", controller.RejectBookIssueRequest).Methods("POST")
	r.HandleFunc("/fetchbookissuerequests", controller.FetchBookIssueRequests).Methods("GET")
	r.HandleFunc("/fetchissuedbooklist", controller.FetchBookIssuedList).Methods("GET")
	http.ListenAndServe("localhost:8000", r)
}
