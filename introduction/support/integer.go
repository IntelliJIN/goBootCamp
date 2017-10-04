// +build OMIT

package main

import "fmt"

func main() {
	// START1 OMIT
	// целые числа
	var defaultInt int64
	intVar := 10
	// STOP1 OMIT
	// START2 OMIT
	fmt.Println("defaultInt = ", defaultInt)
	fmt.Println("2 + 3        = ", 2+3) // 5
	fmt.Println("2 - 3        = ", 2-3) // -1

	intVar--
	fmt.Println("intVar--     = ", intVar) // 9
	intVar++
	fmt.Println("intVar++     = ", intVar) // 10
	// STOP2 OMIT
}
