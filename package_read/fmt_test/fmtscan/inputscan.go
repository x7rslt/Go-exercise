//从标准输入读取内容，赋值给预定义变量

package main

import "fmt"

func main() {
	var (
		a string
		b int
		c float32
		d bool
	)
	fmt.Scan(&a, &b, &c, &d)
	fmt.Println(a, b, c, d)
}
