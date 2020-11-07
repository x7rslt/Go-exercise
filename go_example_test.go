package go_example_test

import (
	"fmt"
	"os"
	"testing"
)

//闭包
func Add(m int) func(n int) int {
	return func(n int) int {
		num := n * m //运算要和后面的return拆分
		return num
	}
}

func TestClosures(t *testing.T) {
	num := Add(10)
	fmt.Println(num(10))

}

//递归Recursion
func Sum(i int) int {
	//fmt.Println(i)
	if i < 1000000 {
		i = i * 10
		fmt.Println(i)
		return Sum(i)
	}
	return 999999
}

func TestRecursion(t *testing.T) {
	fmt.Println(Sum(1))
}

//方法：Go支持在结构类型上定义的方法

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func TestMethod(t *testing.T) {
	var tangle rect = rect{
		width:  12,
		height: 15,
	}
	fmt.Println(tangle.area())
	fmt.Println(tangle.perim())

	rtangle := &tangle

	fmt.Println(rtangle.area())
	fmt.Println(rtangle.perim())

}

func TestArgs(t *testing.T) {
	content := os.Args[1:]
	fmt.Println(content)
}
