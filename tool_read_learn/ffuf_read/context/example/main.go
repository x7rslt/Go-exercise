package main

import (
	"context"
	"fmt"
)

func main() {
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				//这里ctx.Done迷糊很久，本质上就是做一个判断，如果cancel()触发，case<-cancel.Done()
				//执行，goroutine退出。
				case <-ctx.Done():
					return
				case dst <- n:
					n++
				}
			}
		}()
		fmt.Printf("dst type:%T \n", dst)
		fmt.Println(n)
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	fmt.Println("Context Start.")
	defer cancel()

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 20 {
			break
		}
	}
	fmt.Println("Context done.")

}
