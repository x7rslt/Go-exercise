package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	r := gin.Default()
	r.GET("/cookie",func(c *gin.Context){
		cookie,err := c.Cookie("gin_cookie")
		if err!=nil{
			cookie = "NotSet"
			c.SetCookie("gin_cookie","test",3600,"/","localhost",false,true)

		}
		c.String(http.StatusOK,"%v",c.Request.Header)
		fmt.Printf("Cookie value:%s\n",cookie)
	})
	r.Run()
}