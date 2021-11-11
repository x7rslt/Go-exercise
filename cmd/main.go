package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

//not supported by windows
func main() {
	path, err := exec.LookPath("nslookup")
	if err != nil {
		panic(err)
	}
	fmt.Println(path)
	args := []string{"nslookup", "www.baidu.com"}
	env := os.Environ()
	execErr := syscall.Exec(path, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
