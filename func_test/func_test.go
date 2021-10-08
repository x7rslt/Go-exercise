package test

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

//Function types and function values can be used and passed around just like other values
// Map applies op to each element of a.
type Operator func(x float64) float64
func Map(op Operator, a []float64) []float64 {
	res := make([]float64, len(a))
	for i, x := range a {
		res[i] = op(x)
	}
	return res
}

func TestFuncType(t *testing.T) {
	op := math.Abs
	a := []float64{1, -2}
	b := Map(op, a)
	fmt.Println(b) // [1 2]

	c := Map(func(x float64) float64 { return 10 * x }, b)
	fmt.Println(c) // [10, 20]
}

//function literal 函数变量
//func sort.Slice(slice interface{}, less func(i, j int) bool)
func TestFuncSlice(t *testing.T){
	people := []string{"Alice", "Bob", "Dave"}
	sort.Slice(people, func(i, j int) bool {
		return len(people[i]) < len(people[j])
	})
	fmt.Println(people)
	// Output: [Bob Dave Alice]
}

//Closures
// New returns a function Count.
// Count prints the number of times it has been invoked.
func New() (Count func()) {
	n := 0
	return func() {
		n++
		fmt.Println(n)
	}
}

func TestClosures(t *testing.T) {
	f1, f2 := New(), New()
	f1() // 1
	f2() // 1 (different n)
	f1() // 2
	f2() // 2
}


type Pinger struct {
	rec func(s string, i int)
}

func Test(t *testing.T) {
	var p Pinger
	p.rec = func(s string, i int) {
		fmt.Println(s, i)
	}
}
