package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Hello(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"Hello":"World！",
	})
}

func Welcome(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"Welcome":"World!",
	})
}

func main(){
	r := gin.Default()
	v1:= r.Group("v1")
	{
		v1.GET("/",Hello)
	}

	v2 := r.Group("v2")
	{
		v2.GET("/",Welcome)
	}
	r.Run()

}