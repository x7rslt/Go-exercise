package main

import "github.com/gin-gonic/gin"

type Person struct{
	ID string `uri:"id"binding:"required,uuid"`
	Name string `uri:"name"binding:"required"`
}

func main(){
	r := gin.Default()
	r.GET("/:name/:id",func(c *gin.Context){
		var person Person
		if err := c.ShouldBindUri(&person);err!=nil{
			c.JSON(400,gin.H{"msg":err})
			return
		}
		c.JSON(200,gin.H{"name":person.Name,"uuid":person.ID})
	})
	r.Run()
}