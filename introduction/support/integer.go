// +build OMIT

package main

import "fmt"

func main() {
	// START1 OMIT
	var defaultInt int
	intVar := 10
	// STOP1 OMIT
	// START2 OMIT
	fmt.Println("defaultInt = ", defaultInt)
	fmt.Println("2 + 3        = ", 2+3)
	fmt.Println("2 - 3        = ", 2-3)
	fmt.Println("2 / 3        = ", 2/3)
	fmt.Println("2 * 3        = ", 2*3)

	intVar--
	fmt.Println("intVar--     = ", intVar)
	intVar++
	fmt.Println("intVar++     = ", intVar)
	intVar += 3
	fmt.Println("intVar += 3  =", intVar)

	// STOP2 OMIT
}
