package main

import (
	"bytes"
	"fmt"
	"unicode/utf8"
)

func comma(s string) string {
	var buf bytes.Buffer           //同buf := bytes.NewBuffer(nil)
	n := utf8.RuneCountInString(s) //各种字符格式都支持
	for _, v := range s {
		buf.WriteRune(v)
		n--
		if n%3 == 0 && n != 0 {
			buf.WriteByte(',')
		}
	}
	return buf.String()
}

func main() {
	s := "一二三四"
	s2 := "123456789"
	fmt.Println(comma(s))
	fmt.Println(comma(s2))
}
