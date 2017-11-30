package main

import (
	"database/sql"
	"net/http"

	"./LoginModule"
	"./SignupModule"
	"./Utils/ConnectDB"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB

func main() {
	db = ConnectDB.ConnectMySQL()
	router := mux.NewRouter()
	router.HandleFunc("/", homePage)

	router.HandleFunc("/login", func(res http.ResponseWriter, req *http.Request) {
		LoginModule.LoginPage(res, req, db)
	})

	router.HandleFunc("/signup", func(res http.ResponseWriter, req *http.Request) {
		SignupModule.SignupPage(res, req, db)
	})
	http.ListenAndServe(":8080", router)
}

func homePage(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "index.html")
}
