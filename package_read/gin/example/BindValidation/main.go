package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//Binding from JSON
type Login struct{
	User string `form:"user" json:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func main(){
	r := gin.Default()
	r.Static("/","./statics")
	//r.StaticFile("/statics","./statics")
	r.POST("/loginJSON",func(c *gin.Context){
		var json Login
		if err := c.ShouldBindJSON(&json);err !=nil{
			c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
			return
		}
		if json.User!="manu" ||json.Password!="123"{
			c.JSON(http.StatusBadRequest,gin.H{
				"Status":"Unauthorized",
			})
			return
		}
		c.JSON(http.StatusOK,gin.H{"Status":"You are logined in."})
	})

	r.POST("/loginXML",func(c *gin.Context){
		var xml Login
		if err := c.ShouldBindXML(&xml);err!= nil{
			c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
			return
		}
		if xml.User != "manu"||xml.Password!="123"{
			c.JSON(http.StatusUnauthorized,gin.H{
				"Status":"Unauthorized",

			})
			return
		}
		c.JSON(http.StatusOK,gin.H{"Status":"You are logined in."})
	})
	r.POST("/loginForm",func(c *gin.Context){
		var form Login
		if err := c.ShouldBind(&form);err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{"Status":err.Error()})
			return
		}
		if form.User != "manu" ||form.Password!="123"{
			c.JSON(http.StatusUnauthorized,gin.H{
				"Status":"Unauthorized",
			})
			return
		}
		c.JSON(http.StatusOK,gin.H{"Status":"You are logined in."})
	})

	r.Run()

}

