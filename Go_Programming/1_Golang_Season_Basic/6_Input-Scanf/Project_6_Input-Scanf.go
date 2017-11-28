package main

import "fmt"

func main() {
	fmt.Print("Input Your Number : ")
	var input float64
	fmt.Scanf("%f", &input)
	output := input * 10
	fmt.Println(output)
}
