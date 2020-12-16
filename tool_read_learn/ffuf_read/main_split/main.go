package main

import (
	"fmt"
	"os"
)

//main可以分拆成多个文件，运行时：go run main.go test.go 即可
//编译同样：go build main.go test.go
func main() {
	//test()
	fmt.Println(os.UserHomeDir())
}
