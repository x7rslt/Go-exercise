package main

import "fmt"

type multiStringFlag []string

func (m *multiStringFlag) String() string {
	return "lehehe"
}
func (m *multiStringFlag) Set(value string) error {
	*m = append(*m, value)
	return nil
}

func main() {
	fmt.Print("test")
	var t multiStringFlag
	fmt.Print(t)
	fmt.Print(t.String())
	t.Set("hello")
	fmt.Println(t)
}
