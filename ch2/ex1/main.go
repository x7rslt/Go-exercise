//导入tempconv包，转换开尔文温度
package main

import (
	"fmt"

	"Go_exercise/ch2/ex1/tempconv"
)

func main() {
	var temp tempconv.Celsius = 10
	fmt.Println(tempconv.CToK(temp))
}
