package main

import (
	"flag"
	"fmt"
)

var nFlag = flag.Int("n", 1234, "help message for flag n")
var flagvar int

func init() {
	flag.IntVar(&flagvar, "flagname", 1234, "help message for flagname")
}
func main() {

	flag.Parse()
	fmt.Println(*nFlag)
	fmt.Println(flagvar)

}

//flag 待学习
