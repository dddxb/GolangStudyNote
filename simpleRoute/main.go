package main

import (
	"GolangStudy/simpleRoute/http/controller"
	"GolangStudy/simpleRoute/lib"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func init() {
	var err error
	lib.Ms, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", "root", "0825", "127.0.0.1", 3306, "one_draw", "utf8mb4"))
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	r := mux.NewRouter()
	route(r)
	const PORT = "8001"
	log.Fatal(http.ListenAndServe("0.0.0.0:"+PORT, r))
}

func route(r *mux.Router) {
	r.HandleFunc("/", controller.HomeController.Home)
	r.HandleFunc("/hello", controller.HomeController.Hello)

	s := r.PathPrefix("/admin").Subrouter()
	s.HandleFunc("/login", controller.AdminController.Login)
	s.HandleFunc("/register", controller.AdminController.Register)

	s = r.PathPrefix("/query").Subrouter()
	s.HandleFunc("/user/{id}", controller.UserController.View).Methods("POST")
}
