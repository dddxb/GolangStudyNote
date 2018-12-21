package controller

import (
	"fmt"
	"net/http"
)

type HomeController struct {
}

func (HomeController) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "home")
}

func (HomeController) Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

func (HomeController) Query(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("query"))
}
