package main

import (
    "fmt"
    "os"
    "os/exec"
)

func main() {
	cmd := exec.Command("nmap","8.8.8.8")
    cmd.Stderr = os.Stderr
    cmd.Stdout = os.Stdout
    cmd.Stdin = os.Stdin
    if err := cmd.Run(); err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

}