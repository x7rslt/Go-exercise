package deadlock_test

import (
	"fmt"
	"testing"
)

func TestDeadLock(t *testing.T) {
	ch := make(chan int)
	ch <- 1
	fmt.Println(<-ch)
	/*
	fatal error: all goroutines are asleep - deadlock!

	goroutine 1 [chan receive]:
	testing.(*T).Run(0xc000100340, {0x110e8a0, 0x119083d7cb639f}, 0x1116fe8)
	*/
}
//Reason
//no other goroutine has access to the channel or the lock,
//a group of goroutines are waiting for each other and none of them is able to proceed.