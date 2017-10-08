package main

import "fmt"

const (
	one = iota
	two
	_ // пустая константа
	four
)

func main() {
	// START2 OMIT
	fmt.Println(one, two, four)
	// STOP2 OMIT
}
