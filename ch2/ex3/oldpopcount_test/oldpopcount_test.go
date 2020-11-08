package oldpopcount_test

import (
	"fmt"
	"testing"

	"gopl.io/ch2/popcount"
)

func BenchmarkOldPopCount(b *testing.B) {
	var x uint64 = 120000000000
	b.ResetTimer()
	fmt.Println(popcount.PopCount(x))
	b.StopTimer()
}
