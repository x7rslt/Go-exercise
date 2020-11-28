package main

import (
	"Go_exercise/ch2/ex3/popcount"
	"fmt"
	"testing"
)

func BenchmarkPopCount(b *testing.B) {
	var x uint64 = 120000000000
	b.ResetTimer()
	fmt.Println(popcount.PopCount(x))
	b.StopTimer()
}
