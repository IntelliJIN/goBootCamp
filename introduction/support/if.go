package main

import (
	"fmt"
	"os"
)

func main() {
	// START1 OMIT
	trueVar := true
	falseVar := false

	fmt.Print("First:  ")
	if falseVar {
		fmt.Println("if falseVar")
	} else if trueVar {
		fmt.Println("if falseVar -> if trueVar")
	} else {
		fmt.Println("if falseVar -> if trueVar -> else")
	}

	fmt.Print("Second: ")
	if _, err := os.Stat("some unknown file"); err == nil {
		fmt.Println("unknown file exists")
	} else if _, err := os.Stat("./support/if.go"); err == nil {
		fmt.Println("real file exists")
	} else {
		fmt.Println("both files not exist")
	}
	// STOP1 OMIT
}