package ex2_test

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	f := 1e100
	i := int(f)
	fmt.Println(i, int64(f))

	var z float64
	//inta := z
	intb := 1 / z
	x := 1 / z
	fmt.Println(intb == x)
}
