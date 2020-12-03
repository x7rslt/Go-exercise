package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func comma(s string) (string, error) {
	if !strings.Contains(s, ".") {
		return s, fmt.Errorf("%s not a float number.", s)
	}
	if _, err := strconv.ParseFloat(s, 64); err != nil {
		return s, err
	}
	sl := strings.Split(s, ".")
	var s1, s2 string
	switch len(sl) {
	case 1:
		s1 = sl[0]
	case 2:
		s1, s2 = sl[0], sl[1]
	default:
		return s, fmt.Errorf("split error")
	}
	buf1 := bytes.NewBuffer(nil)
	n1 := len(s1)
	for _, v := range s1 {
		buf1.WriteRune(v)
		n1--
		if v == '+' || v == '-' {
			continue
		}
		if n1%3 == 0 && n1 > 0 {
			buf1.WriteByte(',')
		}
	}
	buf2 := bytes.NewBuffer(nil)
	var n2 int
	for _, v := range s2 {
		buf2.WriteRune(v)
		n2++
		if n2%3 == 0 && n2 < len(s2) {
			buf2.WriteByte(',')
		}
	}
	return strings.Join([]string{buf1.String(), buf2.String()}, "."), nil
}

func main() {
	s1 := "123456.1"
	s, err := comma(s1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(s)
	}

}
