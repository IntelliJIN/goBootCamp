package main

import (
	"fmt"
)

func main() {
	// START1 OMIT
	//массивы
	var first [2]bool
	second := [...]int{-1, -2, -3}
	third := [6]uint{1, 2, 3, 4, 5, 6}

	var fourth [2][2]float32
	fifth := [2][2]float32{{1}, {5.4, 7.7}}
	// STOP1 OMIT
	// START2 OMIT
	fmt.Printf("first  len(%d) = %v\n", len(first), first)
	fmt.Printf("second len(%d) = %v\n", len(second), second)
	fmt.Printf("third  len(%d) = %v\n", len(third), third)
	fmt.Printf("fourth len(%d) = %v\n", len(fourth), fourth)
	fmt.Printf("fifth  len(%d) = %v\n\n", len(fifth), fifth)

	fmt.Println("second[2]   = ", second[2])   // second[2]   = -3
	fmt.Println("fifth[1][0] = ", fifth[1][0]) // fifth[1][0] = 5.4
	// STOP2 OMIT
}
