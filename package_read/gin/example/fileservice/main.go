package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
//快速开启文件共享服务
func main(){
	r := gin.Default()
	r.StaticFS("/down",http.Dir("."))
	r.Run()
}
