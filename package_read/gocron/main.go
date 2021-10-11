package main

import (
	"fmt"
	"time"

	"github.com/jasonlvhit/gocron"
)
//制定计划任务，执行三秒自动停止
func jobs(quit <-chan bool) {
	for {
		//cron jobs
		g := gocron.NewScheduler()
		g.Every(1).Second().Do(stuff)

		select {
		case <-quit:
			// here we receive from quit and exit
			// if `g` has started, we may want to `g.Clear()`
			// or the scheduled jobs will continue, we will just not be waiting for them to finish.
			return
		case <-g.Start():
			// here we know all the jobs are done, and go round the loop again
		}
	}
}


func stuff() {
	fmt.Println("doing job")
}

func main() {
	q := make(chan bool)
	go jobs(q)
	time.Sleep(3 * time.Second)

	//to quit the goroutine
	q <- true
	close(q)
	fmt.Println("main")
}

//Out put:执行三秒自动结束
//doing job
//doing job
//main
