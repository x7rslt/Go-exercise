package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type OutputOptions struct {
	DebugLog string
}

func main() {
	var (
		buf    bytes.Buffer
		logger = log.New(&buf, "logger: ", log.Lshortfile)
	)

	logger.Print("Hello, log file!")

	fmt.Print(&buf)
	otp := OutputOptions{"log test"}
	log.Print("test")
	fmt.Println(len(otp.DebugLog))
	if len(otp.DebugLog) != 0 {
		f, err := os.OpenFile(otp.DebugLog, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		fmt.Println(f)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Disabling logging,encountered errors(s):%s\n", err)
			log.SetOutput(ioutil.Discard)
		} else {
			log.SetOutput(f)
			defer f.Close()
		}
	} else {
		//设置log输出
		log.SetOutput(ioutil.Discard)
	}
	fmt.Println(otp)
	//log写入到log日志里
	log.Print("test")
}
