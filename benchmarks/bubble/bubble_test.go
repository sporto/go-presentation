package main

import (
	"testing" //import go package for testing related functionality
)

func Benchmark_bubble(b *testing.B) { //benchmark function starts with "Benchmark" and takes a pointer to type testing.B
	for i := 0; i < b.N; i++ {
		arr := []int{2, 10, 1, 9, 5, 6, 8, 3, 7, 4}
		bubble(arr)
	}
}

// go test -file bubble_test.go -bench=".*"
