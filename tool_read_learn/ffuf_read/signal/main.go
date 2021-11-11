package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//signal for interruptMonitor
func interruptMonitor() {
	sigChan := make(chan os.Signal, 2)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGALRM)
	go func() {
		for range sigChan {
			errmsg := "Caught keyboard interrupt (Ctrl-C)\n"
			sl := <-sigChan
			fmt.Println(errmsg, sl)
			os.Exit(1)
		}
	}()
}

func main() {
	interruptMonitor()
	time.Sleep(10 * time.Second)
}
