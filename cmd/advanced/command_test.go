package command_test

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"testing"
)

func TestBasic(t *testing.T){
	cmd := exec.Command("ls", "-lah")
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

func TestCommandDisplay(t *testing.T){
	cmd := exec.Command("ls", "-lah")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("combined out:\n%s\n", string(out))
}

