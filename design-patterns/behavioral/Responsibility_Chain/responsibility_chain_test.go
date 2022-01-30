package Responsibility_Chain

import (
	"fmt"
	"testing"
)

func TestNewHandler(t *testing.T) {
	laowang := NewHandler("laowang",nil,1)
	laoxiao := NewHandler("laoxiao",laowang,2)

	r := laowang.Handler(1)
	fmt.Println(r)
	r2 := laoxiao.Handler(2)
	fmt.Println(r2)
}