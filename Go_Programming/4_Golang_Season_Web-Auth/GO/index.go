package main

import (
	"database/sql"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "root:root@/mydb")
	checkPanicErr(err)
	defer db.Close()
	err = db.Ping()
	checkPanicErr(err)

	http.HandleFunc("/", homePage)
	http.HandleFunc("/login", loginPage)
	http.HandleFunc("/signup", signupPage)
	http.ListenAndServe(":8080", nil)
}

func homePage(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "../HTML/index.html")
}

func loginPage(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.ServeFile(res, req, "../HTML/login.html")
		return
	}
	username := req.FormValue("username")
	password := req.FormValue("password")
	var dbName string
	var dbPass string

	err := db.QueryRow("SELECT mem_user, mem_pwd FROM member WHERE mem_user = ?", username).Scan(&dbName, &dbPass)
	checkErrLoginRedir(err, res, req)

	err = bcrypt.CompareHashAndPassword([]byte(dbPass), []byte(password))
	checkErrLoginRedir(err, res, req)

	res.Write([]byte("Hello : " + dbName))
}

func signupPage(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.ServeFile(res, req, "../HTML/signup.html")
		return
	}
	username := req.FormValue("username")
	password := req.FormValue("password")
	var user string
	err := db.QueryRow("SELECT mem_user FROM member WHERE mem_user = ?", username).Scan(&user)

	switch {
	case err == sql.ErrNoRows:
		hashedPassword, hashErr := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		checkHTTPErr(hashErr, res)
		db.Exec("INSERT INTO member(mem_user, mem_pwd) VALUES(?, ?)", username, hashedPassword)
		res.Write([]byte("User Create!!"))
		return
	case err != nil:
		http.Error(res, "Server error, unable to create your account", 500)
		return
	default:
		http.Redirect(res, req, "/", 301)
	}
}

func checkHTTPErr(err error, res http.ResponseWriter) {
	if err != nil {
		http.Error(res, "Server error, unable to create your account", 500)
		return
	}
}

func checkPanicErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func checkErrLoginRedir(err error, res http.ResponseWriter, req *http.Request) {
	if err != nil {
		http.Redirect(res, req, "/login", 301)
		return
	}
}
