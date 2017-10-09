// +build OMIT

package main

import (
	"fmt"
	"math"
)

func main() {
	// START1 OMIT
	// Пределы значений целых int*
	fmt.Printf("int8:   [%v, %v]\n", math.MinInt8, math.MaxInt8)
	fmt.Printf("int16:  [%v, %v]\n", math.MinInt16, math.MaxInt16)
	fmt.Printf("int32:  [%v, %v]\n", math.MinInt32, math.MaxInt32)
	fmt.Printf("int64:  [%v, %v]\n", math.MinInt64, math.MaxInt64)

	//Пределы значений целых uint*

	fmt.Printf("uint8:  [0, %v]\n", math.MaxUint8)
	fmt.Printf("uint16: [0, %v]\n", math.MaxUint16)
	fmt.Printf("uint32: [0, %v]\n", math.MaxUint32)
	fmt.Printf("uint64: [0, %v]\n", uint64(math.MaxUint64))

	//Синонимы целых типов
	fmt.Println("byte    - int8")
	fmt.Println("rune    - int32")
	fmt.Println("int     - int32, или int64, в зависимости от ОС")
	fmt.Println("uint    - uint32, или uint64, в зависимости от ОС")
	fmt.Println("uintptr - Беззнаковое целое, пригодное для хранения значения указателя")
	// STOP1 OMIT
}
