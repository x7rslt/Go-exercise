package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	r := gin.Default()
	r.GET("/someJSON",func(c *gin.Context){
		data := map[string]interface{}{
			"lang":"GO language",
			"tag":"<br>",
		}

		c.AsciiJSON(http.StatusOK,data)

	})
	r.Run(":8080")
}
