package time_test

import (
	"fmt"
	"log"
	"testing"
	"time"
)

//Measure a piece of code
func TestExecutionTime(t *testing.T){
	start := time.Now()
	// Code to measure
	duration := time.Since(start)

	// Formatted string, such as "2h3m0.5s" or "4.503μs"
	fmt.Println(duration)

	// Nanoseconds as int64
	fmt.Println(duration.Nanoseconds())
}

func foo() {
	defer duration(track("foo"))
	// Code to measure
}
func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}

func TestFuncTime(t *testing.T){
	foo()
}