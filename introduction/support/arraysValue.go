package main

import (
	"fmt"
)

func main() {
	// START1 OMIT
	//массивы
	third := [6]uint{1, 2, 3, 4, 5, 6}
	// STOP1 OMIT
	// START2 OMIT
	copyOfThird := third
	third[4] = 500
	third[5] = 600
	fmt.Println("third       = ", third)
	fmt.Println("copyOfThird = ", copyOfThird)
	// STOP2 OMIT
}
