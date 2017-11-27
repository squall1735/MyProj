package main

import "fmt"

func main() {
	fmt.Print("Input Your Number : ")
	var input float64
	fmt.Scanf("%f", &input)

	condition := input > 2
	if condition {
		fmt.Println("Worked")
	} else {
		fmt.Println("Not Worked")
	}

	// && 2 ดเงื่อนไข จริงทั้งคู่
	if 6 > 3 && 8 > 5 {
		fmt.Println("Worked")
	} else {
		fmt.Println("Now Worked")
	}

	// || เงื่อนไขอย่างใด เป็นจริง 1 อัน
	if 6 < 3 || 8 < 5 {
		fmt.Println("Worked")
	} else {
		fmt.Println("Now Worked")
	}
}
