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

	copyOfFourth := make([]int, 6)
	copy(copyOfFourth, fourth)
	fourth = []int{1, 2, 3, 4, 5, 6}
	// START1 OMIT
	fourth = append(fourth, copyOfFourth[:3]...)
	// STOP1 OMIT
	// START2 OMIT
	fmt.Printf("fourth len(%d) cap(%d)\n", len(fourth), cap(fourth))

	fourth = append(fourth, 111, 222, 333)
	fmt.Println("fourth = ", fourth)
	fmt.Printf("fourth len(%d) cap(%d)\n", len(fourth), cap(fourth))

	fourth = append(fourth, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Printf("fourth len(%d) cap(%d)\n", len(fourth), cap(fourth))
	// STOP2 OMIT
}
