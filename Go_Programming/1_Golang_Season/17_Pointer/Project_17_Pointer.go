package main

import "fmt"

func main() {
	x := 10
	fmt.Println("Value is", x)
	fmt.Printf("Value is %d\n", x)

	// Address
	fmt.Printf("Address x variable %x\n", &x)

	// Pointer
	var p *int
	p = &x // ชี้ไปยัง Address ที่ x อยู่
	fmt.Printf("Pointer p = %x\n", p)

	*p = 20
	fmt.Printf("value is %d\n", x)
}
