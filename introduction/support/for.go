package main

import (
	"fmt"
)

func main() {
	// START1 OMIT
	first := 10
	second := 10
	// STOP1 OMIT
	// START2 OMIT
	fmt.Println("First: first =", first)
	for {
		fmt.Print(".")
		if first <= 0 {
			break
		}
		first--
	}
	fmt.Println("\nfirst =", first)

	fmt.Println("Second: second =", second)
	for second = 10; second < 20; second++ {
		fmt.Print(".")
	}
	fmt.Println("\nsecond =", second)
	// STOP2 OMIT
}
