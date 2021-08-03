package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(Cors())
	r.GET("/test",func(c *gin.Context){
		c.JSON(http.StatusOK,gin.H{
			"Title":"This is a test page.",
		})
	})
	r.Run()


}


func Cors()gin.HandlerFunc{
	return func(c *gin.Context){
		method:= c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization,Token,X-TOKEN ")
		c.Header("Access-Control-Allow-Methods", "POST, GET,PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		if method=="options"{
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()

	}
}
