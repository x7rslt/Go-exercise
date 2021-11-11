package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	r:= gin.Default()
	r.GET("/someDataFromReader",func(c *gin.Context){
		res,err := http.Get("https://img.alicdn.com/tfs/TB1w8Ozl7cx_u4jSZFlXXXnUFXa-2880-120.jpg")
		if err!= nil||res.StatusCode!= http.StatusOK{
			c.Status(http.StatusServiceUnavailable)
			return
		}
		reader := res.Body
		contentLength := res.ContentLength
		contentType:= res.Header.Get("Content-Type")
		extraHeaders := map[string]string{
			"Content-Disposition":`attachment;filename="gopher.png"`,
		}
		c.DataFromReader(http.StatusOK,contentLength,contentType,reader,extraHeaders)
	})
	r.Run()
}
