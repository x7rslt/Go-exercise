package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	reader := strings.NewReader("io.reader test")

	p := make([]byte, 4)

	for {
		n, err := reader.Read(p)
		if err != nil {
			if err == io.EOF {
				fmt.Println("EoF")
				break
			}
			fmt.Println(err)
			os.Exit(1)

		}
		fmt.Println(n, string(p))
	}

	t := log.Writer()
	fmt.Println(t)
}
