package main

import "github.com/gin-gonic/gin"

func main(){
	r := gin.Default()
	r.StaticFile("/static/shellocde.txt","./static/shellocde.txt")
	r.StaticFile("/favicon.ico","./resources/favicon.ico")
	r.Run(":8081")
}
