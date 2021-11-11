//逐行读取.csv文件

package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	//准备读取文件
	fileName := `C:\Users\x7rsl\Desktop\BugHunt\top_websites.csv`
	fs, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("can not open the file, err is %+v", err)
	}
	defer fs.Close()

	newpath := `C:\Users\x7rsl\Desktop\BugHunt\top_websites.txt`
	f, _ := os.Create(newpath)
	defer f.Close()

	r := csv.NewReader(fs)
	//针对大文件，一行一行的读取文件
	for {
		row, err := r.Read()
		if err != nil && err != io.EOF {
			log.Fatalf("can not read, err is %+v", err)
		}
		if err == io.EOF {
			break
		}
		fmt.Println(row[1])
		//d2 := []byte{row[1]}
		f.WriteString(string(`http://`) + row[1] + string('\n'))
	}
}
