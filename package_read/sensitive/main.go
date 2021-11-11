package main

import (
	"fmt"
	"github.com/importcjj/sensitive"
)

func main() {
	filter := sensitive.New()
	err :=filter.LoadWordDict("./dict/dict.txt")
	if err != nil {
		fmt.Println("fail to load dict %v", err)
	}
	// Do something
	re,word := filter.FindIn("隐私数据泄漏")
	fmt.Println(re,word)
}