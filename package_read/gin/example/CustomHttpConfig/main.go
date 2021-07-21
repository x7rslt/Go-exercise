package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func testpage(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"Hello":"world.",
	})
}
func main(){
	r := gin.Default()
	r.GET("/",func(c *gin.Context){
		c.String(200,"Hello")
	})
	r.GET("/test",testpage)
	s:= &http.Server{
		Addr:":8080",
		Handler:r,
		ReadTimeout: 10*time.Second,
		WriteTimeout: 10*time.Second,
		MaxHeaderBytes: 1<<20,
	}

	s.ListenAndServe()

}
