package main

import (
	"fmt"
)

func main() {
	// START1 OMIT
	fmt.Print("First:  ")
	switch {
	case 1 < 2:
		fmt.Println("1 < 2") // First:  TRUE  1 < 2
	case 2 > 3:
		fmt.Println("2 > 3")
	}

	fmt.Print("Second: ")
	switch suffix := "jpg"; suffix {
	case "png":
		fallthrough
	case "jpg":
		fallthrough
	case "gif":
		fmt.Println("Image") // Second: Image
	default:
		fmt.Println("Not image")
	}
	// STOP1 OMIT
}