package main

import "fmt"

func main() {
	x := make([]float64, 5)
	fmt.Println(x)

	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	slice2 := append(slice1, 9, 10)
	fmt.Println(slice2)

	arr := [5]float64{1, 2, 3, 4, 5}
	y := arr[0:4]
	fmt.Println(y)

	slice11 := []int{1, 2, 3}
	slice22 := make([]int, 2)
	copy(slice22, slice11)
	fmt.Println(slice22)
}
