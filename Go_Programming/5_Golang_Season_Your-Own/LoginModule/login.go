package LoginModule

import (
	"database/sql"
	"net/http"

	"../Bean"
	"../Utils/CheckErr"

	"golang.org/x/crypto/bcrypt"
)

// Member ...
// type Member struct {
// 	memUser string
// 	memPwd  string
// }

// LoginPage ..
func LoginPage(res http.ResponseWriter, req *http.Request, db *sql.DB) {
	if req.Method != "POST" {
		http.ServeFile(res, req, "./LoginModule/login.html")
	}
	// var mem Member
	// mem.memUser = req.FormValue("username")
	// mem.memPwd = req.FormValue("password")

	var member Bean.Member
	member.MemUser = req.FormValue("username")
	member.MemPwd = req.FormValue("password")
	// username := req.FormValue("username")
	// password := req.FormValue("password")

	var dbName string
	var dbPass string

	err := db.QueryRow("SELECT mem_user, mem_pwd FROM member WHERE mem_user = ?", member.MemUser).Scan(&dbName, &dbPass)
	CheckErr.CheckRedirErr(err)

	err = bcrypt.CompareHashAndPassword([]byte(dbPass), []byte(member.MemPwd))
	CheckErr.CheckRedirErr(err)

	res.Write([]byte("Hello : " + dbName))
}
