package main

import (
	"fmt"

	// "../../FUNCTION"	// Upper Folder
	"./FUNCTION" // Sub Folder
)

func main() {
	c := FUNCTION.CallBack()
	fmt.Println("c : " + c)
}
