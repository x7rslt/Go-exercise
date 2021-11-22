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

func main(){
	funcresult := expand("hello foo",func (s string)string{
		return "Bye " + s
	})
	fmt.Println(funcresult)

}