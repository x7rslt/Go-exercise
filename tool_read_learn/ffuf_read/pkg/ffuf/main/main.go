package main

import "fmt"

type multiStringFlag []string

func (m *multiStringFlag) String() string {
	return "lehehe"
}

func main() {
	fmt.Print("test")
	var t multiStringFlag
	fmt.Print(t.String())
}
