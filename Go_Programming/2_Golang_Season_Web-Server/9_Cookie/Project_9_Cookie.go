package main

import (
	"fmt"
	"net/http"
	"time"
)

// Cookie is Simple-Data to use
type Cookie struct {
	Name    string
	Value   string
	Expires time.Time
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	exp := time.Now().Add(time.Hour * 24 * 36)
	cookie := http.Cookie{Name: "user", Value: "thanakorn", Expires: exp}
	http.SetCookie(w, &cookie)
	fmt.Fprintf(w, "Create Cookie")
}
