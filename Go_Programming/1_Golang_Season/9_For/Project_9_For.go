package main

import "fmt"

func main() {
	// แบบที่ 1
	for i := 1; i <= 10; i++ {
		fmt.Println("THANAKORN")
	}

	// แบบที่ 2
	j := 1
	for j <= 10 {
		fmt.Println("THANAKORN")
		// j++
		j = j + 1
	}

	// Odd - Even
	for i := 1; i <= 10; i++ {
		if (i % 2) == 0 {
			fmt.Println(i, " : Even")
		} else {
			fmt.Println(i, " : Odd")
		}
	}
}
