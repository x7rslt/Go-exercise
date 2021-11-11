package main

import (
	"flag"
	"fmt"
	"time"
)

var period = flag.Duration("period", 3*time.Second, "Sleep period.")

func main() {
	flag.Parse()
	fmt.Printf("Sleeping for %v...", *period)
	time.Sleep(*period)
	fmt.Println()
}
