package main

import (
	"math/rand"
	"testing"
)

func Add2(a, b int) int {
	return a + b
}

func GenerateRandomSlice(size int) []int {
	slice := make([]int, size)

	for i := range slice {
		slice[i] = rand.Intn(100)
	}

	return slice
}

func SumSlice(slice []int) int {
	sum := 0
	for _, v := range slice {
		sum += v
	}

	return sum
}

func TestGenerateRandomSlice(t *testing.T) {
	size := 100
	slice := GenerateRandomSlice(size)

	if len(slice) != size {
		t.Errorf("Expected slice size %d, received %d", size, len(slice))
	}
}

func BenchmarkGenerateRandomSlice(b *testing.B) {
	for range b.N {
		GenerateRandomSlice(1000)
	}
}

func BenchmarkSumSlice(b *testing.B) {
	slice := GenerateRandomSlice(1000)
	b.ResetTimer() // exclude the setup time

	for range b.N {
		SumSlice(slice)
	}
}

/*
go test -bench=. -memprofile mem.prof benchmarking_test.go|grep -v 'cpu:'
BenchmarkGenerateRandomSlice-10           226928              5338 ns/op
BenchmarkSumSlice-10                     5287311               248.7 ns/op

- RandomSlice was executing 226,928 times during the benchmark and each execution took 5,338 nanoseconds / 0.5338ms.
- SumSlice was executing 5,287,311 times during the benchmark and each execution took 248.7 nanoseconds / 0.02487ms.
- It will generate mem.prof file

Run: go tool pprof
*/

// go test -bench=. -benchmem benchmarking_test.go|grep -v 'cpu:'

// func BenchmarkAddSmallInput(b *testing.B) {
// 	for range b.N {
// 		Add2(2, 3)
// 	}
// }

// func BenchmarkAddMediumInput(b *testing.B) {
// 	for range b.N {
// 		Add2(200, 300)
// 	}
// }

// func BenchmarkAddLargeInput(b *testing.B) {
// 	for range b.N {
// 		Add2(2000, 3000)
// 	}
// }
