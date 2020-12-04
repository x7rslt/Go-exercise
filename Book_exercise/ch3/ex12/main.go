//判断同文异构字符串

package main

import (
	"fmt"
	"strings"
)

func judge(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for _, r := range s1 {
		rs := string(r)
		if strings.Count(s1, rs) != strings.Count(s2, rs) {
			return false
		}
	}
	return true
}

func main() {
	s1 := "str"
	s2 := "rts"
	fmt.Println(judge(s1, s2))
}
