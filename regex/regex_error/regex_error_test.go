package regex_error_test

import (
	"fmt"
	"regexp"
	"testing"
)

//Golang 的 regexp 对网络流量做正则匹配时，发现有些情况无法正确进行匹配，找到资料发现 regexp 内部以 UTF-8 编码的方式来处理正则表达式，而网络流量是字节序列，由其中的非 UTF-8 字符造成的问题。
func TestError(t *testing.T){
	src := "hello你好\xfftest"
	regexp1 := "hello"
	fmt.Println(regexp.MatchString(regexp1,src))
	regexp2 := "你"
	fmt.Println(regexp.MatchString(regexp2,src))
	regexp3 := "\xff"
	fmt.Println(regexp.MatchString(regexp3,src))

}