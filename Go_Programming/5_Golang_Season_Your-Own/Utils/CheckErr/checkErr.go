package CheckErr

import (
	"net/http"
)

// CheckHTTPErr ...
func CheckHTTPErr(err error, res http.ResponseWriter) {
	if err != nil {
		http.Error(res, "Server error, unable to create your account", 500)
		return
	}
}

// CheckPanicErr ...
func CheckPanicErr(err error) {
	if err != nil {
		// fmt.Println("panic(err.Error())")
		panic(err.Error())
	}
}

// CheckRedirErr ...
// func CheckRedirErr(err error, res http.ResponseWriter, req *http.Request) {
func CheckRedirErr(err error) {
	if err != nil {
		// http.Redirect(res, req, "/login", 301)
		http.RedirectHandler("/login", 301)
		return
	}
}
