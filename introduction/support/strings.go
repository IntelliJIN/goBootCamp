// +build OMIT

package main

import "fmt"

func main() {
	// START1 OMIT
	var hello string = "Hello\n\t"
	var hello2 string = `Hello\n\t`
	var world = "World"
	var world2 = `World`
	var defaultString string
	// STOP1 OMIT
	// START2 OMIT
	fmt.Println(hello, world)
	fmt.Println(hello2, world2)
	fmt.Println("defaultString: ", defaultString)
	// STOP2 OMIT
}
