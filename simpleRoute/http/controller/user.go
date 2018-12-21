package controller

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type UserController struct {
}

func (UserController) View(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "Your input id is: %s", vars["id"])
}
