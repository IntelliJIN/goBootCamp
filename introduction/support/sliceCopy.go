package main

import (
	"fmt"
)

func main() {
	fourth := []int{1, 2, 3, 4, 5, 6}

	part1 := fourth[:2]
	part2 := fourth[2:4]
	part3 := fourth[4:]

	part1[0] = 100
	part2[0] = 300
	part3[0] = 500
	// START1 OMIT
	copyOfFourth := make([]int, 6)
	count := copy(copyOfFourth, fourth)
	fourth = []int{1, 2, 3, 4, 5, 6}
	// STOP1 OMIT
	// START2 OMIT
	fmt.Println("fourth        = ", fourth)
	fmt.Println("copyOfFourth  = ", copyOfFourth)
	fmt.Println("count         = ", count)
	// STOP2 OMIT
}
