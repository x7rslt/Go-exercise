//练习1.11:使用更长的参数列表来尝试fetchall，例如使用alexa.com排名前100万的网站，如果一个网站没有响应，程序的行为是怎样？（8.9节会通过复制这个例子来描述响应机制）
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
	//"strings"
	"encoding/csv"
    "log"
)

func main(){
	start := time.Now()
	ch := make(chan string)
	fileName := "/Users/xiaoshuai/go/src/exercise/top.csv"
	fs, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("can not open the file, err is %+v", err)
	}
	defer fs.Close()
	
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
		go fetch("http://"+row[1],ch)
	}
	for range ch{
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elasped\n",time.Since(start).Seconds())
}

func fetch(url string,ch chan<- string){
	start := time.Now()
	resp,err := http.Get(url)
	if err != nil{
		//ch <- fmt.Sprint(err)
		ch <- "Error"
		return
	}
	//filename := "./" + strings.Split(url,".")[1]
	//content,err := os.Create(filename)
	//num,err := io.Copy(content,resp.Body)
	resp.Body.Close()
	//fmt.Println("This url response :",num)
	if err !=nil{
		//ch<- fmt.Sprint(err)
		ch <- "Error--2"
		return
	}
	fmt.Println(resp.Status)
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %s",secs,url)
}