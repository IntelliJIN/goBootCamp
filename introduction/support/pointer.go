package main

import (
	"fmt"
)

func main() {
	// START1 OMIT
	intVar := 10

	firstPointer := &intVar
	secondPointer := &firstPointer
	var thirdPointer *int
	// STOP1 OMIT
	tableRowFormat := "| %-15v | %-8T | %-12p | %-12v |\n"

	fmt.Println("| Название        | Тип      | Адрес        | Значение     |")
	fmt.Println("|-----------------|----------|--------------|--------------|")
	fmt.Printf(tableRowFormat, "intvar", intVar, &intVar, intVar)
	fmt.Println("|-----------------|----------|--------------|--------------|")
	fmt.Printf(tableRowFormat, "firstPointer", firstPointer, &firstPointer, firstPointer)
	fmt.Printf(tableRowFormat, "*firstPointer", *firstPointer, &*firstPointer, *firstPointer)
	fmt.Println("|-----------------|----------|--------------|--------------|")
	fmt.Printf(tableRowFormat, "secondPointer", secondPointer, &secondPointer, secondPointer)
	fmt.Printf(tableRowFormat, "*secondPointer", *secondPointer, &*secondPointer, *secondPointer)
	fmt.Printf(tableRowFormat, "**secondPointer", **secondPointer, &**secondPointer, **secondPointer)
	fmt.Println("|-----------------|----------|--------------|--------------|")
	fmt.Printf("| %-15v | %-8T | %-12p | %-12p |\n", "thirdPointer", thirdPointer, &thirdPointer, thirdPointer)
	// START2 OMIT
	// STOP2 OMIT
}
