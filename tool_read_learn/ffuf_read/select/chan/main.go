package main

import (
	"fmt"
)

//rename_test.go test wordlist Set()
func main() {
	fmt.Println("tset")
	ch := make(chan int, 1)
	go func() {
		ch <- 10
		//fmt.Println(<-ch) 在这里没法接收，要在goroutine外面接收
	}()
	fmt.Println(<-ch)

}
