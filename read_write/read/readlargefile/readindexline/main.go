package main

import (
	"fmt"
	"os"
)
//读取超大文件的指定行
const filename = "filename.txt"

func main() {
	readFile(filename)
}

func readFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	buf := make([]byte, 32)
	stat, err := os.Stat(filename)
	fmt.Println("Stat:",stat)
	start := stat.Size() - 32
	fmt.Println("File size:",stat.Size())
	_, err = file.ReadAt(buf, start)
	if err == nil {
		fmt.Printf("%s\n", buf)
	}

}