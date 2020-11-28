//练习1.7: 函数io.Copy(dst,src)从src读，写入dst。使用它代替ioutil.ReadAll来复制响应到os.Stdout,这样不需要装下整个响应数据流到缓冲区。确保检查io.Copy返回到错误结果

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main(){
	for _,url := range os.Args[1:]{
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