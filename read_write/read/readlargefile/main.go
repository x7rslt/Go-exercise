package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)
//大文件逐行读取，搜索关键字
func main(){
	//准备读取文件
	fileName := "filename.txt"
	fs, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("can not open the file, err is %+v", err)
	}
	defer fs.Close()

	if err != nil{
		fmt.Println("Open file Err:",err)
	}
	defer fs.Close()

	reader := bufio.NewReader(fs)
	t1 := time.Now()
	fmt.Println("File keyword search:")
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Printf("%#v\n", line)
				break
			}
			panic(err)
		}
		if strings.Contains(line,"15555001522"){
			fmt.Println(line)
			fmt.Println(time.Since(t1))
		}
		//fmt.Printf("%#v\n", line)
	}
	/*
	//容易因为单行长度过长报错：bufio.Scanner: token too long；改用上面的方式；解析参考：https://github.com/ma6174/blog/issues/10
	scanner := bufio.NewScanner(fs)
	t1 := time.Now()
	fmt.Println("File keyword search:")
	for scanner.Scan(){
		if strings.Contains(scanner.Text(),"15867187152"){
			fmt.Println(scanner.Text())
			fmt.Println(time.Since(t1))
		}
	}
	if err := scanner.Err();err !=nil{
		fmt.Println("Line read err:",err)
	}
*/

}