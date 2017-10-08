package main

import (
	"fmt"
)

func main() {
	// START1 OMIT
	third := []string{"a", "b", "c", "d", "e", "f"}
	fourth := map[string]string{"a": "A", "b": "B", "c": "C", "d": "D", "e": "E", "f": "F"}
	// STOP1 OMIT
	// START2 OMIT
	fmt.Println("Third: third =", third)
	for index := range third {
		fmt.Println(index, " -> ", third[index])
	}

	fmt.Println("Fourth: fourth =", fourth)
	for key, value := range fourth {
		fmt.Println(key, " -> ", value)
	}
	// STOP2 OMIT
}
