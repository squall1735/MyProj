package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// ResultData is Data from Query Table: product
type ResultData struct {
	ID    int
	Name  string
	Price string // ไม่ได้เอาไป + - * /
}

var templates = template.Must(template.ParseFiles("index.html"))

func main() {
	http.HandleFunc("/", result)
	http.HandleFunc("/delete", delete)
	http.ListenAndServe(":8080", nil)
}

func result(res http.ResponseWriter, req *http.Request) {
	var db, err = sql.Open("mysql", "root:root@/mydb")
	checkFatalErr(err)
	defer db.Close()
	rows, err := db.Query("SELECT product_id, product_name, product_price FROM product")
	checkPanicErr(err)

	tRes := ResultData{}
	var results []ResultData

	for rows.Next() {
		var id int
		var name, price string

		err = rows.Scan(&id, &name, &price)
		tRes.ID = id
		tRes.Name = name
		tRes.Price = price

		results = append(results, tRes)
		checkPanicErr(err)
	}
	templates.Execute(res, results)
	fmt.Println(results)
}

func delete(res http.ResponseWriter, req *http.Request) {
	var db, err = sql.Open("mysql", "root:root@/mydb")
	checkFatalErr(err)
	stmt, err := db.Prepare("DELETE FROM product WHERE product_id=?")
	stmt.Exec(req.URL.Query().Get("id"))
	checkPanicErr(err)
	http.Redirect(res, req, "/", 301)
}

func checkFatalErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func checkPanicErr(err error) {
	if err != nil {
		panic(err)
	}
}
