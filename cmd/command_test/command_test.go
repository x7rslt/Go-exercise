package command_test

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"syscall"
	"testing"
)


//执行shell文件
func TestShell(t *testing.T){
	path,_ := os.Getwd()
	fmt.Println(path)
	result,err := exec.Command("date").Output()  //output可以显示命令执行结果和报错
	if err != nil{
		log.Fatal(err)
	}
	fmt.Printf("This day is %s",result)
	out,err := exec.Command("/bin/sh","sample.sh").Output()
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(out)
}

//执行多个命令
func TestMultiCommand(t *testing.T){
	var ip []byte
	var err error
	var cmd *exec.Cmd
	result,err := exec.Command("/bin/sh","-c","ls;date;echo 'Hello world'").Output()
	if err !=nil{
		log.Fatal(err)
	}
	fmt.Println("Command execute result:",string(result))

	// 执行连续的shell命令时, 需要注意指定执行路径和参数, 否则运行出错
	cmd = exec.Command("/bin/sh", "-c", `/sbin/ifconfig en0 | grep -E 'inet ' |  awk '{print $2}'`)
	if ip, err = cmd.Output(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(ip))
	fmt.Println(strings.Trim(string(ip), "\n"))
}

//run a command and show output
func TestOut1(t *testing.T){
	cmd := exec.Command("ls", "-lah")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}
//效果同上
func TestOut2(t *testing.T){
	cmd := exec.Command("ls", "-lah")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("combined out:\n%s\n", string(out))

}
//分别输出结果和报错，很少用到
func TestStdoutStderr(t *testing.T){
	cmd := exec.Command("ls", "-lah")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	fmt.Printf("out:\n%s\nerr:\n%s\n", outStr, errStr)
}

//显示输出和命令执行过程
func TestOutputProgress(t *testing.T){
	cmd := exec.Command("ls", "-lah")
	if runtime.GOOS == "windows" {
		cmd = exec.Command("tasklist")
	}

	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)
	cmd.Stderr = io.MultiWriter(os.Stderr, &stderrBuf)

	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	outStr, errStr := string(stdoutBuf.Bytes()), string(stderrBuf.Bytes())
	fmt.Printf("\nout:\n%s\nerr:\n%s\n", outStr, errStr)
}

func TestCmdRun(t *testing.T){
	cmd := exec.Command("ping", "127.0.0.1")  //自动等待命令执行完成
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	//err = cmd.Wait() //The Wait method will return the exit code and release associated resources once the command exits.
	log.Printf("Command finished with error: %v", err)
}
func TestCmdStart(t *testing.T){
	cmd := exec.Command("ping", "127.0.0.1")  //默认不等待命令执行完成，需要在下面添加wait等待
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	err = cmd.Wait() //The Wait method will return the exit code and release associated resources once the command exits.
	log.Printf("Command finished with error: %v", err)
}

//Terminat process start with os.exec
func TestKill(t *testing.T){
	cmd := exec.Command("ping", "www.baidu.com")
	cmd.Start()
	fmt.Println(cmd.Process.Pid) // 打印进程pid
	cmd.Process.Kill() // 你要的结束进程
	cmd.Process.Signal(syscall.SIGINT) // 给进程发送信号
	cmd.Wait()
}