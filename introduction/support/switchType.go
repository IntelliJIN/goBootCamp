package main

import (
	"fmt"
)

func main() {
	// START1 OMIT
	var someObject interface{}
	someObject = 5

	fmt.Print("Third:  ")
	switch someObject.(type) {
	case bool:
		fmt.Println("bool")
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("Sorry, I don't know") // Third:  Sorry, I don't know
	}
	// STOP1 OMIT
}