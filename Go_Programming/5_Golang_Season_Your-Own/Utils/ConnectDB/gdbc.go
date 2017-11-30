package ConnectDB

import (
	"database/sql"

	"../CheckErr"
)

// ConnectMySQL ...
func ConnectMySQL() *sql.DB {
	var db *sql.DB
	var err error
	db, err = sql.Open("mysql", "root:root@/mydb")

	// If Error DB is Close.
	CheckErr.CheckPanicErr(err)
	// defer db.Close()

	// Check Status(ping) Database
	err = db.Ping()
	CheckErr.CheckPanicErr(err)

	return db
}
