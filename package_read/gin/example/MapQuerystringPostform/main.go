package main


import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	r := gin.Default()
	r.POST("/post",func(c *gin.Context){
		ids := c.QueryMap("ids")
		names:= c.PostFormMap("names")
		fmt.Println("ids : %v;names:%v",ids,names)
		c.String(http.StatusOK,"ids : %v;names:%v",ids,names)
	})
	r.Run()

}

/*POST /post?ids[a]=1234&ids[b]=hello HTTP/1.1
Content-Type: application/x-www-form-urlencoded

names[first]=thinkerou&names[second]=tianou*/