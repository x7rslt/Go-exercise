package main_test

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"testing"
)

func TestExec(t *testing.T) {
	d, err := exec.LookPath("ping")
	fmt.Printf("Search command :%v ; Result:%v\n", d, err)

	d2, err := exec.LookPath("dir")
	fmt.Printf("Search command :%v ; Result:%v\n", d2, err)

	cmd := exec.Command("ping")
	cmd.Stdin = strings.NewReader("127.0.0.1")
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	fmt.Printf("in all caps :%v\n", out.String())

	out2, err := exec.Command("whoami").Output()
	fmt.Printf("%s", out2) //out2's type is unint8 , %s format print it's string.

	path, err := exec.LookPath("systeminfo")
	if err != nil {
		log.Fatal("installing system is in your future")
	}
	fmt.Println("Path:", path)
}
