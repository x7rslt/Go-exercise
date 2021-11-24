package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
//快速开启文件共享服务
func main(){
	r := gin.Default()
	//file,_:= filepath.Abs("./")
	//fmt.Println("Current path : ",file)
	r.StaticFS("/down",http.Dir("."))
	r.Run()
}
