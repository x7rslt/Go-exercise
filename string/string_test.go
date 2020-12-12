package string_test

import (
	"fmt"
	"log"
	"testing"
)

func Test(t *testing.T) {
	str := "Go 爱好者"
	log.Printf("The string:%q\n", str)
	fmt.Printf("=>runes(char):%q\n", []rune(str))
	fmt.Printf("=>runes(hex):%x\n", []rune(str))
	fmt.Printf("=>bytes(hex):[% x]\n", []byte(str))
	fmt.Printf("str 有%d个字节\n", len(str))
	fmt.Printf("str 有%d个字符\n", len([]rune(str)))
	for i, j := range str {
		fmt.Println(i, j)
		fmt.Printf("%d,%q\n", i, j)
	}
}
