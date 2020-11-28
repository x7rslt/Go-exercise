package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"Go_exercise/ch2/ex3/popcount"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	fmt.Println("Please input something:")
	for input.Scan() {
		s := input.Text()
		if s == "q" {
			fmt.Println("Programm exit!")
			os.Exit(0)
		}
		x, err := strconv.ParseUint(s, 10, 64)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Please input a number!", err)
			continue
		}
		fmt.Println(popcount.PopCount(x))
	}
}
