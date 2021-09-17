package iterator_test

import (
	"fmt"
	"testing"
)


//Basic iterator pattern
func Iterate(f func(n int)){
	for i:= 1;i<=3;i++{
		f(i)
	}
}



func TestIterator(t *testing.T){
	Iterate(func(n int){
		fmt.Println(n)
	})
}


//Iterator with break
// Iterate calls the f function with n = 1, 2, and 3.
// If f returns true, Iterate returns immediately
// skipping any remaining values.
func Iterate2(f func(n int) (skip bool)) {
	for i := 1; i <= 3; i++ {
		if f(i) {
			return
		}
	}
}


func TestIteratorBreak(t *testing.T){
	Iterate2(func(n int) (skip bool) {
		fmt.Println(n)
		return n == 2
	})
}