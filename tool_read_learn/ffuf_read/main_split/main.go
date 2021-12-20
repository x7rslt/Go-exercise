package main

import (
	"fmt"
	"os"
)

//main可以分拆成多个文件，运行时：go run rename_test.go test.go 即可
//编译同样：go build rename_test.go test.go
func main() {
	test()
	fmt.Println(os.UserHomeDir())
}
