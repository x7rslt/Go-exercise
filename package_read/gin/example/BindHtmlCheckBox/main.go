package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type myForm struct{
	Colors []string `form:"colors[]"`
}

func formHandler(c *gin.Context){
	var fakeForm myForm

	c.ShouldBind(&fakeForm)
	c.JSON(http.StatusOK,gin.H{
		"Color":fakeForm.Colors,
	})
}
func indexHandler(c *gin.Context){
	c.HTML(200,"form.html",nil)
}

func main(){
	r := gin.Default()
	r.LoadHTMLGlob("view/*")
	r.GET("/",indexHandler)
	r.POST("/",formHandler)
	r.Run(":8080")
}
