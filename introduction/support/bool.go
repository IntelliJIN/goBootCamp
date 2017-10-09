// +build OMIT

package main

import "fmt"

func main() {
	// START1 OMIT
	//булевые переменные
	var first, second bool
	var third bool = true
	fourth := !third
	var fifth = true
	// STOP1 OMIT
	// START2 OMIT
	fmt.Println("first second: ", first, second)
	fmt.Println("third: ", third)
	fmt.Println("fourth: ", fourth)
	fmt.Println("fifth: ", fifth)

	fmt.Println("!true  = ", !true)

	fmt.Println("true && true   = ", true && true)         // true
	fmt.Println("true && false  = ", true && false)        // false
	fmt.Println("false && false = ", false && false, "\n") // false

	fmt.Println("true || true   = ", true || true)         // true
	fmt.Println("true || false  = ", true || false)        // true
	fmt.Println("false || false = ", false || false, "\n") // false

	fmt.Println("2 < 3  = ", 2 < 3) // true
	fmt.Println("2 > 3  = ", 2 > 3) // false
	// STOP2 OMIT
}
