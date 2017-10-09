// +build OMIT
package main

import (
	"fmt"
)

func main() {
	//срезы
	fourth := []int{1, 2, 3, 4, 5, 6}

	part1 := fourth[:2]
	part2 := fourth[2:4]
	part3 := fourth[4:]
	part4 := fourth[:]
	// START1 OMIT
	part1[0] = 100
	part2[0] = 300
	part3[0] = 500
	// STOP1 OMIT
	// START2 OMIT
	fmt.Println("fourth = ", fourth)
	fmt.Println("part4  = ", part4)
	// STOP2 OMIT
}
