package main

import "fmt"

// Keyword : tys
type Golang struct {
	title    string
	author   string
	subtitle string
	price    float64
}

func main() {
	var book Golang
	book.title = "Go Programming"
	book.author = "THANAKORN"
	book.subtitle = "Thanakorn Tutorial"
	book.price = 199.99

	fmt.Println(book)
	fmt.Println(book.price)

	book1 := Golang{title: "Go Programming", author: "THANAKORN", price: 199.99}
	fmt.Println(book1.price * 2)
}
