package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[:])
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	flag.ErrHelp = errors.New("flag:help requested.")
	var Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}

	Usage()
}
