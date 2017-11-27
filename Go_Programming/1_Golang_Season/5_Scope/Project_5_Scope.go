package main

import "fmt"

// SCOPE
// GLOBAL VARIABLE
var gVariable int = 500

func main() {
	lVariable := 40

	fmt.Println("Global =", gVariable)
	fmt.Println("Local =", lVariable)
	anotherFunction()
}

func anotherFunction() {
	lVariable := 20

	fmt.Println("Global =", gVariable)
	fmt.Println("Local =", lVariable)
}
