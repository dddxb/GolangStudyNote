package main

import (
	"ginStart/lib"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	Db := lib.SQL{}
	Db.Connect("mysql", "127.0.0.1", 3306, "root", "0825", "one_draw", "utf8mb4")
	r := mux.NewRouter()
	lib.Route(r)
	const PORT = "8888"
	http.ListenAndServe("0.0.0.0:"+PORT, r)
}
