package main

import (
	"log"
	"os"
	"os/exec"
	"runtime"
)
//执行多个命令

func main(){
	cmd := exec.Command("/bin/sh","-c","subfinder -d douban.com -all | httpx")
	if runtime.GOOS == "windows" {
		cmd = exec.Command("tasklist")
	}
	//使Go执行系统命令显示结果
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}