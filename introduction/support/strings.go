// +build OMIT

package main

import "fmt"

func main() {
	// START1 OMIT
	//строки
	var hello string = "Hello\n\t"
	var world = "World"
	var defaultString string
	// STOP1 OMIT
	// START2 OMIT
	fmt.Println(hello, world)
	fmt.Println("defaultString: ", defaultString)
	// STOP2 OMIT
}
