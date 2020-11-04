package pointer_test

import (
	"fmt"
	"testing"
)

func zeroval(i int) {
	i = 0
}

func zeroptr(i *int) {
	*i = 0
}

func TestPointer(t *testing.T) {
	n := 1
	fmt.Println(&n)
	fmt.Println(n)
	zeroval(n)
	fmt.Println(n)

	fmt.Println("---------------")
	zeroptr(&n)
	fmt.Println(n)
	fmt.Println(&n)
}
