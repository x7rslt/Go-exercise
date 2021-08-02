package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	r := gin.Default()
	r.GET("/json",func(c *gin.Context){
		c.JSON(200,gin.H{
			"html":"<b>Hello world!</b>",
		})
	})

	r.GET("/purejson",func(c *gin.Context){
		c.PureJSON(http.StatusOK,gin.H{
				"Html":"<b>Hello, PureJson test.</b>",
		})
	})
	r.Run()
}
