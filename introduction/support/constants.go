package main

import "fmt"

const (
	a = 1
	b
	c = 2
	d
)

func main() {
	// START2 OMIT
	fmt.Println(a, b, c, d)
	// STOP2 OMIT
}
