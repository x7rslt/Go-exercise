package idgener

import (
	"fmt"
	"testing"
)

func BenchmarkIGener(b *testing.B) {
	ig := NewIGener()
	for n := 0; n <= b.N; n++ {
		<-ig
		//fmt.Println(<-ig)
	}
}

func TestIGener(t *testing.T) {
	ig := NewIGener()
	var id string
	m := make(map[string]int)
	for i := 0; i < 10000; i++ {
		id = <-ig
		m[id] = i
	}
	fmt.Print(m)
	if len(m) != 10000 {
		fmt.Println("TestIGener Error")
	}
	fmt.Println("Test IGener PASS")
}
