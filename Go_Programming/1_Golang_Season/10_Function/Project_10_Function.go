package main

import "fmt"

func main() {
	doSomething("THANAKORN")
	doSomething("TUTORIAL")
	addition(8, 2)
	result := addition2(10, 10)
	fmt.Println(result * 10)
}

// ไม่รับ ไม่คืน
func doSome() {
	fmt.Println("OK")
}

// รับค่า
func doSomething(str string) {
	fmt.Println(str)
}

func addition(a int, b int) {
	fmt.Println(a + b)
}

// คืนค่า
func addition2(a int, b int) int {
	output := a + b
	return output
}
