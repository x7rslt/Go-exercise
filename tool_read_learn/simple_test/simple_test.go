package simple

import (
	"fmt"
	"strings"
	"testing"
)

//main.go test wordlist Set()
func Test(t *testing.T) {
	fmt.Println("lehehe")
	value := "simple,test"
	delimited := strings.Split(value, ",")
	fmt.Println(delimited)

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
				//fmt.Println(<-ch)
			}
		}
	}()
	//fmt.Println("Last:", <-ch)

}
