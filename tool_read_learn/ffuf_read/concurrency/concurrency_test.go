package concurrency_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
func TestGoroutine(t *testing.T) {
	say("Hello") //如果主程序执行完，goroutine无法执行
	go say("world.")

	say("Hello 2:") //如果主程序继续执行，goroutine参与执行

}

//Channels
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func TestChannel(t *testing.T) {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}

//Buffered Channels
func TestBufferChannel(t *testing.T) {
	c := make(chan int, 2)
	c <- 1
	c <- 2
	fmt.Println(<-c)
	fmt.Println(<-c)
	c <- 3
	fmt.Println(<-c)
}

//Range and close
/* A sender can close a channel to indicate that no more values will be sent. Receivers can test whether
a channel has been closed by assigning a second parameter to the receive expression:after
	v, ok:= <-ch
ok is false if there are no more values to receive and the channel is closed.
The loop for i:= range c recieves values from the channel repeatedly until it is closed.
Note:Only the sender should close a channel, never the receiver. Sending
on a closed channel will cause a panic.
Another note:Channel aren't like files;you don't usually need to close them.
Closing is only necessary when the receiver must be told there are
no more values coming, such as to terminate a range loop.

*/
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func TestRangeChannel(t *testing.T) {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}

}

//sync.Mutex
/*We've seen how channels are great for communicate among
goroutines.
But what if we don't need communication?What if we just want to
make sure only one goroutine can access a variable at a time to
avoid conflicts?
This concept is called mutual exclusion, and the conventional name
for the data structure that provides it is mutex.
Go's standard library provides mutual exclusion with sync.Mutex and its
two methods.
Lock
Unlock
*/

type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	c.v[key]++
	c.mu.Unlock()
}

func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	//Lock so only one goroutine at a time can access the map c.v.
	defer c.mu.Unlock()
	return c.v[key]
}

func TestMutex(t *testing.T) {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}
	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
