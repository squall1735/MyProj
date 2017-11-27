package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/login", login)
	http.HandleFunc("/login-event", loginEvent)
	http.HandleFunc("/upload", upload)
	http.HandleFunc("/upload-event", uploadHandle)
	http.ListenAndServe(":8080", nil)
}

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

func upload(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../HTML/upload.html")
}

func uploadHandle(w http.ResponseWriter, r *http.Request) {
	// file: data, handle: details of file, err: message error
	file, handle, err := r.FormFile("file")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	// handle => Map
	fmt.Fprintf(w, "%v", handle.Header)
	// func Create: https://golang.org/pkg/os/
	f, err := os.OpenFile("../UPLOAD/"+handle.Filename, os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
	fmt.Fprintf(w, "Upload complete")
}
