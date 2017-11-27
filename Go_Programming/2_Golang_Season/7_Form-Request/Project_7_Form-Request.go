package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

// Product is a representation of a product
type Product struct {
	Name  string
	Price int
}

func main() {
	var templates = template.Must(template.ParseFiles("../HTML/index.html"))
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		myProduct := Product{"นมสด", 500}
		templates.ExecuteTemplate(w, "index.html", myProduct)
	})

	http.ListenAndServe(":8080", router)
}
