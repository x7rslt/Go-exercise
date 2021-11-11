package select_test

import (
	"fmt"
	"testing"
	"time"
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

//Select default:
//The default case in a select is run no other case is ready.
// Use a default case to try a send or receive without blocking.

func TestSelectDefault(t *testing.T) {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return //通道使用return终端后续
		default:
			fmt.Println("   .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

//Select exercise
func PT(a, b chan string) {
	text := "Sunshine."
	for {
		select {
		case a <- text:
			fmt.Println("Get something.")
		case end := <-b:
			fmt.Println(end, "Quit.")
			return
		}
	}
}

func TestS(t *testing.T) {
	a := make(chan string)
	b := make(chan string)
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println(<-a)
		}
		b <- "Quit time."
	}()
	PT(a, b)
}
