package fib

import (
	"testing"
	//"bytes"
)

var result int

//func BenchmarkFib10(b *testing.B) {
//	// run the Fib function b.N times
//	for n := 0; n < b.N; n++ {
//		Fib(10)
//	}
//}

//
//func BenchmarkFib10(b *testing.B) {
//	var r int
//	for n := 0; n < b.N; n++ {
//		// always record the result of Fib to prevent
//		// the compiler eliminating the function call.
//		r = Fib(10)
//	}
//	// always store the result to a package level variable
//	// so the compiler cannot eliminate the Benchmark itself.
//	result = r
//}

func BenchmarkFib10(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var r int
			r = Fib(10)
			result = r
		}

	})
}
