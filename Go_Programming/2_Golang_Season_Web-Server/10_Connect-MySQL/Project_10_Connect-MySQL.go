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

	rows, err := db.Query("SELECT product_id, product_name, product_price FROM product ORDER BY product_id ASC")
	checkErr(err)

	for rows.Next() {
		var productID int
		var productName string
		var productPrice float64
		err = rows.Scan(&productID, &productName, &productPrice)
		checkErr(err)
		fmt.Printf("productId: %d productName: %s productPrice: %.2f\n", productID, productName, productPrice)
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
