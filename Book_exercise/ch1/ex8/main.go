//练习1.8: 修改fetch程序添加一个http://前缀（假如URL参数缺失协议前缀）。可能会用到strings.HasPrefix

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main(){
	for _,url := range os.Args[1:]{
		if !strings.HasPrefix(url,"http"){
			url = "http://" + url
		}
		resp,err := http.Get(url)
		if err != nil{
			fmt.Fprintf(os.Stderr,"fetch:%v\n",err)
		}
		num,err := io.Copy(os.Stdout,resp.Body)
		resp.Body.Close()
		fmt.Println("This url response :",num)
		if err !=nil{
			fmt.Fprintf(os.Stderr,"fetch:reading %s:%v\n",url,err)
			os.Exit(1)
		}
	}
}