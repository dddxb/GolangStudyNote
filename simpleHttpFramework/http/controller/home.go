package controller

import (
	"fmt"
	"net/http"
)

type HomeControllerBase struct {
}

var HomeController HomeControllerBase

func (HomeControllerBase) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "home")
}

func (HomeControllerBase) Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

func (HomeControllerBase) Query(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("query"))
}
