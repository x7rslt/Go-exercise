package main

import (
	"fmt"
	"strings"
)


func expand(s string, f func(string) string) string{
	var content string
	for _,i := range strings.Split(s," "){
		if i =="foo"{
			content = f(i)
		}
	}
	return content
}

func Expand(s string, f func(string) string) string {
	ret := strings.Replace(s, "$foo", f("foo"), 1024)
	//fmt.Println(f("foo"))
	return ret
}

func main(){
	funcresult := expand("hello foo",func (s string)string{
		return "Bye " + s
	})
	fmt.Println(funcresult)

	funcReplace := Expand("$foo Foo Test. $foo",func(string)string{
		return "Good, This is correct!"
	})
	fmt.Println(funcReplace)
}