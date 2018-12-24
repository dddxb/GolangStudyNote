package http

import (
	"GolangStudyNote/simpleHttpFramework/http/controller"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type RouteBase struct {
}

var Route = RouteBase{}

func (RouteBase) Register(r *mux.Router) {
	r.Use(loggingMiddleware)

	r.HandleFunc("/", controller.HomeController.Home)
	r.HandleFunc("/hello", controller.HomeController.Hello)

	s := r.PathPrefix("/admin").Subrouter()
	s.HandleFunc("/login", controller.AdminController.Login)
	s.HandleFunc("/register", controller.AdminController.Register)

	s = r.PathPrefix("/query").Subrouter()
	s.HandleFunc("/user/{id}", controller.UserController.View).Methods("POST")
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf(fmt.Sprintf("url: %s  method:%s\n", r.RequestURI, r.Method))
		next.ServeHTTP(w, r)
	})
}
