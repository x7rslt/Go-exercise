package popcount_test

import (
	"fmt"
	"testing"

	popcount2 "Go_exercise/ch2/ex4/popcount"

	popcount1 "gopl.io/ch2/popcount"
)

func BenchmarkOldPopCount1(b *testing.B) {
	var x uint64 = 12
	b.ResetTimer()
	fmt.Println(popcount1.PopCount(x))
	b.StopTimer()
}
func BenchmarkOldPopCount2(b *testing.B) {
	var x uint64 = 12
	b.ResetTimer()
	fmt.Println(popcount2.PopCount(x))
	b.StopTimer()
}
