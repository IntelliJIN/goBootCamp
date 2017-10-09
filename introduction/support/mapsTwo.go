package main

import (
	"fmt"
)

func main() {
	first := make(map[bool]float32, 20)
	third := map[string]uint{"one": 1, "two": 2, "three": 3}
	// START1 OMIT
	first[false] = 100
	first[true] = 200

	third["four"] = 4
	third["five"] = 5
	// STOP1 OMIT
	// START2 OMIT
	fmt.Println("first  =", first)
	fmt.Println("third  =", third)

	delete(third, "three")
	delete(third, "two")

	fmt.Println("third =", third)
	value, ok := third["four"]
	if ok {
		fmt.Println("Value: ", value)
	}
	// STOP2 OMIT
}
