package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	r := gin.Default()
	r.Static("/assets","./assets")
	r.StaticFS("/more_statics",http.Dir("./"))
	r.StaticFile("/favicon.ico","./resources/favicon.ico")
	r.Run()
}