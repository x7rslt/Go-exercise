package main

import "fmt"

const (
	KB = 1e3
	MB = 1e6
	GB = 1e9
	TB = 1e12
	PB = 1e15
	EB = 1e18
	ZB = 1e21
	YB = 1e24
)

func main() {
	fmt.Printf("%T, %[1]g\n", KB)
	fmt.Printf("%T, %[1]g\n", MB)
	fmt.Printf("%T, %[1]g\n", GB)
	fmt.Printf("%T, %[1]g\n", TB)
	fmt.Printf("%T, %[1]g\n", PB)
	fmt.Printf("%T, %[1]g\n", EB)
	fmt.Printf("%T, %[1]g\n", ZB)
	fmt.Printf("%T, %[1]g\n", YB)
}
