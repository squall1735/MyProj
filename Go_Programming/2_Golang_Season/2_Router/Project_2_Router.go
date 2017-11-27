package main

import (
	"fmt"
	"net/http"
)

func main() {
	userDB := map[string]int{
		"thanakorn": 20,
		"golang":    30,
		"java":      40,
		"php":       50,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Thanakorn")
	})

	// Ex: http://localhost:8080/user/php
	http.HandleFunc("/user/", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Path[len("/user/"):]
		age := userDB[name]
		fmt.Fprintf(w, "My Name is %s %d year old", name, age)
	})

	http.ListenAndServe(":8080", nil)
}
