package main

import (
  "fmt"
  "net/http"

  "github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *Request)  {
  // Keyword: fp
  fmt.Println(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main()  {
  // Keyword: hf
  http.HandleFunc("/", handler)
  // Keyword: las
  http.ListenAndServe(":8080", nil)
}
