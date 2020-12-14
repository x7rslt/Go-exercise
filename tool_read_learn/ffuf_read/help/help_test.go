package help_test

import (
	"fmt"
	"testing"
)

func Pr(h string) {
	if h == "lehehe" {
		return
	}
	fmt.Println("hahaha")
}
func Test(t *testing.T) {
	h := "leheh"
	Pr(h)
	max_length := 1
	format := fmt.Sprintf("  -%%-%ds %%s", max_length)
	fmt.Println(format)
	fmt.Printf(format, "lehehe", "Test")

	fmt.Printf("\n-%s %s\n", "test", "lehehe")
	fmt.Printf("-%-3s %s\n", "test", "lhehehe")
}

//-1s lehehe%!(EXTRA string=Test)--- PASS: Test (0.00s)
