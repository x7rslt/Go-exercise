//练习1.9: 修改fetch程序来输出HTTP的状态码，可以在resp.Status中找到它

package main

import (
	"fmt"
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
		//num,err := io.Copy(os.Stdout,resp.Body)
		resp.Body.Close()
		//fmt.Println("This url response :",num)
		fmt.Println("Http status:",resp.Status)
	}
}