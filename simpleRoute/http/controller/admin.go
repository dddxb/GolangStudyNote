package controller

import (
	"net/http"
)

type AdminController struct {

}

func (AdminController) Login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("admin/login"))
}

func (AdminController) Register(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("admin/register"))
}