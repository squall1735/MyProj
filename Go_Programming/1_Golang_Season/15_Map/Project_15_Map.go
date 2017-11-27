package main

import "fmt"

func main() {
	x := [5]float64{1, 2, 3, 4, 5}
	fmt.Println(x[0])

	// Map
	y := make(map[string]string)
	y["TH"] = "THAILAND"
	y["JP"] = "JAPAN"
	y["EN"] = "ENGLAND"
	fmt.Println(y["TH"])

	z := map[string]string{
		"TH": "THAILAND",
		"JP": "JAPAN",
	}
	fmt.Println(z)
}
