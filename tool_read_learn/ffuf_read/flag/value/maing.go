package main

import (
	"flag"
	"fmt"
)

type hello struct {
	hi string
}

func (v hello) String() string {
	return "hello string"
}

func (v hello) Set(s string) error {
	v.hi = s
	return nil
}

var u = "test"

func main() {
	fs := flag.NewFlagSet("ExampleValue", flag.ExitOnError)
	fs.Var(&hello{u}, "hello", "How are you")
	fs.Parse([]string{"-hello", "this is command"})
	fmt.Println(u)
	fs.Usage()
}
