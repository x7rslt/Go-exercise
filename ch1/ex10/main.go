//练习1.10: 找一个产生大量数据的网站，连续两次运行fetchall，看报告的时间是否有大的变化，调查缓存情况。
//每一次获取的内容一样吗？修改fetchall将内容输出到文件，这样检查它们是否一致
//go run main.go http://www.baidu.com http://www.tmall.com http://www.jd.com http://www.taobao.com 执行命令
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
	"strings"
)

func main(){
	start := time.Now()
	ch := make(chan string)
	for _,url := range os.Args[1:]{
		go fetch(url,ch)
	}
	for range os.Args[1:]{
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elasped\n",time.Since(start).Seconds())
}

func fetch(url string,ch chan<- string){
	start := time.Now()
	resp,err := http.Get(url)
	if err != nil{
		ch <- fmt.Sprint(err)
		return
	}
	filename := "./" + strings.Split(url,".")[1]
	content,err := os.Create(filename)
	num,err := io.Copy(content,resp.Body)
	resp.Body.Close()
	fmt.Println("This url response :",num)
	if err !=nil{
		ch<- fmt.Sprint(err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s",secs,num,url)
}