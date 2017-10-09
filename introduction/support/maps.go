package main

import (
	"fmt"
)

func main() {
	// START1 OMIT
	first := make(map[bool]float32, 20)

	second := make(map[string]bool)

	third := map[string]uint{"one": 1, "two": 2, "three": 3}

	fourth := map[string]uint{}
	// STOP1 OMIT
	// START2 OMIT
	fmt.Println("first  =", first)
	fmt.Println("second =", second)
	fmt.Println("third  =", third)
	fmt.Println("fourth =", fourth)
	// STOP2 OMIT
}
