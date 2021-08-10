package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func middle()gin.HandlerFunc{
	return func(c *gin.Context){
		log.Println("I am middle 1")
		c.Next()
		log.Println("I am middle 1")
	}
}

func middle2()gin.HandlerFunc{
	return func(c *gin.Context) {
		log.Println("I am middle 2")
		c.Next()
		log.Println("I am middle 2")
	}
}
func main(){
	r := gin.Default()
	r.GET("/test",func(c *gin.Context){
		c.JSON(http.StatusOK,gin.H{
			"Hello":"World",
		})
	})


	v1 := r.Group("/v1")
	v1.Use(middle()).Use(middle2())
	v1.GET("/group",func(c *gin.Context){
		log.Println("Group test.")
		c.JSON(http.StatusOK,gin.H{
			"Group":"GetTets",
		})
	})
	r.Run()
}
