package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Person struct{
	Goods string `form:"goods"`
	Prices int `form:"price"`
}

func startPage(c *gin.Context){
	var person Person
	if c.ShouldBind(&person)==nil{
		log.Println("Goods",person.Goods)
		log.Println("Price",person.Prices)
	}
	c.JSON(http.StatusOK,gin.H{
		"Goods":person.Goods,
		"Prices":person.Prices,
		"Status":"Success",
	})
}
func main(){
	r := gin.Default()
	r.GET("/bill",startPage)
	r.Run()
}