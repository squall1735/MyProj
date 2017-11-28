package main

import "fmt"

func main() {
	// การดำเนินการ ทางคณิตศาสตร์
	fNumber1 := 50
	fNumber2 := 5

	fmt.Println(fNumber1 + fNumber2)
	fmt.Println(fNumber1 - fNumber2)
	fmt.Println(fNumber1 * fNumber2)
	fmt.Println(fNumber1 / fNumber2)

	// ชุดข้อความ
	p1 := "Thanakorn"
	p2 := "Tutorial"
	// Concat
	p3 := p1 + p2
	fmt.Println(p3)

	fmt.Println(p3[0:3])
	// T - 0  => T h a => 3

	fmt.Println(p3[3:])
	// T - 0  => T h a n a k o r n

	// ASCII
	fmt.Println(p3[0])
}
