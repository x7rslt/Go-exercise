package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	r := gin.Default()

	//This handler will match /user/jhon will not match /user/ or /user
	r.GET("/user/:name",func(c *gin.Context){
		name := c.Param("name")
		c.String(http.StatusOK,"Hello %s",name)

	})


	//Match /user/john/ and /user/john/send
	//if no other routers match /user/jhon wil redirect to /user/jhon/
	r.GET("/user/:name/*action",func(c *gin.Context){
		name:= c.Param("name")
		action := c.Param("action")
		message := name + "is" + action
		c.String(http.StatusOK,message)

	})
	r.Run()
}
