package main

import (
	"fmt"
)

func main() {
	// START1 OMIT
	//срезы
	fourth := []int{1, 2, 3, 4, 5, 6}
	part1 := fourth[:2]
	part2 := fourth[2:4]
	part3 := fourth[4:]
	part4 := fourth[:]
	// STOP1 OMIT
	// START2 OMIT
	fmt.Printf("part1 len(%d) cap(%d) = %v\n", len(part1), cap(part1), part1)
	fmt.Printf("part2 len(%d) cap(%d) = %v\n", len(part2), cap(part2), part2)
	fmt.Printf("part3 len(%d) cap(%d) = %v\n", len(part3), cap(part3), part3)
	fmt.Printf("part4 len(%d) cap(%d) = %v\n", len(part4), cap(part4), part4)
	//STOP2 OMIT
}
