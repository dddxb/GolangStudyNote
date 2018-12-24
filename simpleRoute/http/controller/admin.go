package controller

import (
	"GolangStudy/simpleRoute/lib"
	"fmt"
	"log"
	"net/http"
)

type AdminControllerBase struct {
}

var AdminController AdminControllerBase

func (AdminControllerBase) Login(w http.ResponseWriter, r *http.Request) {
	prepare, err := lib.Ms.Prepare(`insert into user (nickname,avatar,openid) values(?,?,?)`)
	if err != nil {
		log.Fatal(err)
	}
	res, err := prepare.Exec("pacno966", "asdhnashdasd", "sadhuashdsad")
	if err != nil {
		w.WriteHeader(200)
		panic(err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(id)

	w.Write([]byte("admidddn/login"))
}

func (AdminControllerBase) Register(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("admin/register"))
}
