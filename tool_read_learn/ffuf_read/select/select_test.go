package select_test

import (
	"fmt"
	"testing"
)

//The select statement lets a goroutine wait on multipe communciation
//operations.

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
func Test(t *testing.T) {
	//A select blocks until one of its cases can run, then
	// it executes that case. It chooses one at random if multiple
	// are ready.
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)

}
