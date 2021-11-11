package main

import (
	"fmt"
	"syscall"
    "os/exec"
)

//启动一个应用，并隐藏Go cmd窗口
func main() {
	cmd_path := "C:\\Windows\\system32\\cmd.exe"
	cmd_instance := exec.Command(cmd_path, "/c", "notepad")
	cmd_instance.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmd_output, err := cmd_instance.Output()
	fmt.Println(cmd_output,err)
}