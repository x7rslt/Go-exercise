package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	r := gin.Default()
	handlers := func(c *gin.Context) {
		data := map[string]interface{}{
			"foo": "bar",
		}
		c.JSONP(http.StatusOK,data)
	}
	r.GET("/JSONP", handlers)
	r.Run()
}


//curl http://127.0.0.1:8080/JSONP?callback=x