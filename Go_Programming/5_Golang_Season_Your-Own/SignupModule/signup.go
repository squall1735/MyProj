package SignupModule

import (
	"database/sql"
	"net/http"

	"../Utils/CheckErr"
	"golang.org/x/crypto/bcrypt"
)

// SignupPage ...
func SignupPage(res http.ResponseWriter, req *http.Request, db *sql.DB) {
	if req.Method != "POST" {
		http.ServeFile(res, req, "./SignupModule/signup.html")
		return
	}
	username := req.FormValue("username")
	password := req.FormValue("password")
	var user string
	err := db.QueryRow("SELECT mem_user FROM member WHERE mem_user = ?", username).Scan(&user)

	switch {
	case err == sql.ErrNoRows:
		hashedPassword, hashErr := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		CheckErr.CheckHTTPErr(hashErr, res)
		db.Exec("INSERT INTO member(mem_user, mem_pwd) VALUES(?, ?)", username, hashedPassword)
		res.Write([]byte("User Create!!"))
		return
	case err != nil:
		http.Error(res, "Server error, unable to create your account", 500)
		return
	default:
		http.Redirect(res, req, "/", 301)
	}
}
