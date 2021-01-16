package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type OutputOptions struct {
	DebugLog string
}

func main() {
	otp := OutputOptions{"log test"}
	fmt.Println(len(otp.DebugLog))
	if len(otp.DebugLog) != 0 {
		f, err := os.OpenFile(otp.DebugLog, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		fmt.Println(f)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Disabling logging,encountered errors(s):%s\n", err)
			log.SetOutput(f)
		} else {
			log.SetOutput(f)
			defer f.Close()
		}
	} else {
		log.SetOutput(ioutil.Discard)
	}
	fmt.Println(otp)
}
