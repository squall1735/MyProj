package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/login", login)
	http.HandleFunc("/login-event", loginEvent)
	http.ListenAndServe(":8080", nil)
}

// Change index to login.html => Quick Tutorial
func index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../HTML/index.html")
}

func login(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../HTML/login.html")
}

func loginEvent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method:", r.Method)
	r.ParseForm()
	fmt.Println("Username:", r.Form["username"])
	fmt.Println("Password:", r.Form["password"])
}
