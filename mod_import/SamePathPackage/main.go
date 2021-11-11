//go mod模式，倒入本地自定义包；go run main.go正常运行；注意：vs插件code runner快捷运行可能会报错：无法找到package。

package main

import (
	"fmt"

	"Go_exercise/mod_import/SamePathPackage/tempconv"
)

func main() {
	var t tempconv.Celsius
	t = 100
	f := tempconv.Fahrenheit(t)
	fmt.Println(f)
}
