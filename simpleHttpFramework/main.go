package main

import (
	http2 "GolangStudyNote/simpleHttpFramework/http"
	"GolangStudyNote/simpleHttpFramework/lib"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func init() {
	lib.Initialize()
}

func main() {
	r := mux.NewRouter()
	http2.Route.Register(r)
	fmt.Printf("Server is run at %s:%s\n", os.Getenv("HTTP_ADDRESS"), os.Getenv("HTTP_PORT"))
	log.Fatal(http.ListenAndServe(os.Getenv("HTTP_ADDRESS")+":"+os.Getenv("HTTP_PORT"), r))
}