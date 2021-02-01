package main

import (
	"bufio"
	"fmt"
	"os"
)

func WL() {
	path := "/Users/xiaoshuai/go/src/Go_exercise/tool_read_learn/ffuf_read/main/input/wordlist/target.txt"
	file, _ := os.Open(path)
	defer file.Close()

	var data [][]byte
	reader := bufio.NewScanner(file)
	for reader.Scan() {
		text := reader.Text()
		fmt.Println(text)
		data = append(data, []byte(text))
	}
	fmt.Println(data)
}
func main() {
	WL()
}
