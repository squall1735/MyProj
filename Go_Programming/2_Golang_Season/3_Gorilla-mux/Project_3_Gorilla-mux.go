package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	userDB := map[string]int{
		"thanakorn": 20,
		"golang":    30,
		"java":      40,
		"php":       50,
	}

	// Installation(HAVE GOROOT): go get github.com/gorilla/mux
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Thanakorn")
	})

	router.HandleFunc("/user/{name}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		name := vars["name"]
		age := userDB[name]
		fmt.Fprintf(w, "My Name is %s %d year old", name, age)
	}).Methods("GET")

	http.ListenAndServe(":8080", router)
}
