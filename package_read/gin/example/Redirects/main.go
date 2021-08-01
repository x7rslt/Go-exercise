package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	r := gin.Default()
	r.GET("/test",func(c *gin.Context){
		c.Request.URL.Path= "/test2"
		r.HandleContext(c)

	})
	r.GET("/test2",func(c *gin.Context){
		c.JSON(http.StatusOK,gin.H{
			"Hello":"world",
		})
	})
	r.GET("/redirects",func(c *gin.Context){
		c.Redirect(http.StatusFound,"https://www.baidu.com")
	})
	r.Run()
}