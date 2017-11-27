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
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO product (product_name, product_price, image_fk, brand_fk) VALUE(?, ?, ?, ?)")
	if err != nil {
		fmt.Println(err)
	}
	stmt.Exec("Test Prod-name", 500.50, 1, 1)
}
