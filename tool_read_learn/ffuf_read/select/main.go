package main

import (
	"fmt"
)

func main() {

	ch := make(chan int, 1)
	go func() {
		for i := 0; i < 10; i++ {
			select {
			case <-ch:
				//"<-ch"在这里只做情景判断，不算接收
				//fmt.Println(<-ch) 算接受，循环结束
				fmt.Println("To conitue")
			case ch <- i: //第一步：i传入ch
				fmt.Println("ch", i)

			}
		}
	}()
	fmt.Println("test")
	fmt.Println(<-ch) //如果没有这里的接收，程序无输出
	//没搞懂为什么和simple_test的情况相同
}
