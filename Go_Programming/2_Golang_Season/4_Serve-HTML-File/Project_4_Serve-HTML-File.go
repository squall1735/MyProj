package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../HTML/index.html")
	})

	router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../HTML/login.html")
	})

	router.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../HTML/signup.html")
	})

	router.HandleFunc("/File", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../name.txt")
	})

	http.ListenAndServe(":8080", router)
}
