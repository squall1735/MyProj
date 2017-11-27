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

	rows, err := db.Query("SELECT product_pk, product_name, product_price FROM product")
	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		var productPk int
		var productName string
		var productPrice float64
		err = rows.Scan(&productPk, &productName, &productPrice)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("productPk: %d productName: %s productPrice: %.2f\n", productPk, productName, productPrice)
	}
}
