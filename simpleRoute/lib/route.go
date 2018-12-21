package lib

import (
	"ginStart/http/controller"
	"github.com/gorilla/mux"
)

func Route(r *mux.Router) {
	r.HandleFunc("/", controller.HomeController{}.Home)
	r.HandleFunc("/hello", controller.HomeController{}.Hello)

	s := r.PathPrefix("/admin").Subrouter()
	s.HandleFunc("/login", controller.AdminController{}.Login)
	s.HandleFunc("/register", controller.AdminController{}.Register)

	s = r.PathPrefix("/query").Subrouter()
	s.HandleFunc("/user/{id}", controller.UserController{}.View).Methods("POST")
}
