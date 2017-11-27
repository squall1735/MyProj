package main

import "fmt"

func main() {
	summation(9, 10)
}

func summation(args ...int) {
	var total int
	for _, n := range args {
		total += n
	}
	fmt.Println(total)
}
