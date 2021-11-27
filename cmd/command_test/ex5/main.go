package main

import (
	"log"
	"os/exec"
)

func main() {
	//gowitness single http://www.baidu.org 多个参数
	cmd := exec.Command("gowitness","single", "https://readhub.cn")

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}
}
