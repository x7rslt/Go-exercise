//导入tempconv包，转换开尔文温度
package main

import (
	"fmt"

	"tempconv"
)

func main() {
	var temp tempconv.Celsius = 10
	fmt.Println(tempconv.CToK(temp))
}
