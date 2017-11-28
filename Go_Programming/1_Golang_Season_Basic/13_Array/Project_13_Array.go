package main

import "fmt"

func main() {
	var x [5]int
	x[0] = 10
	x[1] = 50
	x[2] = 25
	x[3] = 30
	x[4] = 100
	fmt.Println(x)
	fmt.Println(x[3])

	y := [5]float64{10, 20, 30, 40, 50}
	fmt.Println(y)
	var total float64
	for _, value := range y {
		total += value
	}
	fmt.Println(total)
	fmt.Println(total / float64(len(x)))
}
