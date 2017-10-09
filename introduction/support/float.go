// +build OMIT

package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// START1 OMIT
	var defaultFloat float32
	floatVar := 5.5
	// STOP1 OMIT
	// START2 OMIT
	fmt.Println("defaultfloat       = ", defaultFloat)
	fmt.Printf("floatVar (%s) = %v\n", reflect.TypeOf(floatVar), floatVar)

	fmt.Println("MAX float32        = ", math.MaxFloat32)
	fmt.Println("MIN float32        = ", math.SmallestNonzeroFloat32, "\n")

	fmt.Println("MAX float64        = ", math.MaxFloat64)
	fmt.Println("MIN float64        = ", math.SmallestNonzeroFloat64, "\n")
	// STOP2 OMIT
}
