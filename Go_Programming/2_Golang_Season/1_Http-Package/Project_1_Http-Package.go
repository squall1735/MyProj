package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Thanakorn")
	})

	http.HandleFunc("/product", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Product Request")
	})

	http.HandleFunc("/product-function", product)
	http.HandleFunc("/user-function", user)

	http.ListenAndServe(":8080", nil)
}

func product(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Product Function Request")
}

func user(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "User Function Request")
}
