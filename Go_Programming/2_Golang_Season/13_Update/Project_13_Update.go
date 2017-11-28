package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Connect DB
	user := "root"
	pass := "root"
	dbname := "mydb"
	// db, err := sql.Open("mysql", "user:password@/dbname")
	db, err := sql.Open("mysql", user+":"+pass+"@/"+dbname)
	checkErr(err)
	defer db.Close()

	// stmt, err := db.Prepare("UPDATE product SET product_name = ? WHERE product_id = ?")
	stmt, err := db.Prepare("UPDATE product SET product_price = ? WHERE product_id > ?")
	// stmt.Exec("Mouse", 5)
	stmt.Exec(99.9, 4)
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
