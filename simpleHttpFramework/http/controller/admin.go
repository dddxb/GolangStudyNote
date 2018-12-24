package controller

import (
	"GolangStudyNote/simpleHttpFramework/lib"
	"fmt"
	"net/http"
)

type AdminControllerBase struct {
}

var AdminController AdminControllerBase

func (AdminControllerBase) Login(w http.ResponseWriter, r *http.Request) {
	openid := r.FormValue("openid")
	prepare, err := lib.Db.Prepare(`insert into user (nickname,avatar,openid) values(?,?,?)`)
	if err != nil {
		lib.Sugar.Infow(err.Error())
		w.Write([]byte(err.Error()))
		return
	}
	res, err := prepare.Exec("pacno966", "asdhnashdasd", openid)
	if err != nil {
		lib.Sugar.Infow(err.Error())
		w.Write([]byte(err.Error()))
		return
	}
	id, err := res.LastInsertId()
	if err != nil {
		lib.Sugar.Infow(err.Error())
		w.Write([]byte(err.Error()))
		return
	}
	fmt.Fprintf(w, "Insert user success, id is %d!", id)
}

func (AdminControllerBase) Register(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("admin/register"))
}
