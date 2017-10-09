package main

import (
	"fmt"
)

func main() {
	// START1 OMIT
	//срезы
	first := make([]byte, 5, 15)
	second := make([]byte, 3)
	var third []bool
	fourth := []int{1, 2, 3, 4, 5, 6}
	// STOP1 OMIT
	// START2 OMIT
	fmt.Printf("first  len(%d) cap(%d) = %v\n", len(first), cap(first), first)
	fmt.Printf("second len(%d) cap(%d) = %v\n", len(second), cap(second), second)
	fmt.Printf("third  len(%d) cap(%d) = %v\n", len(third), cap(third), third)
	fmt.Printf("fourth len(%d) cap(%d) = %v\n", len(fourth), cap(fourth), fourth)
	//STOP2 OMIT
}
