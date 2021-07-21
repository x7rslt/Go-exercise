package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)
func main(){
	r := gin.New()
	r.Use(gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		return fmt.Sprintf(
			params.ClientIP,
			params.TimeStamp.Format(time.RFC1123),
			params.Method,
			params.Path,
			params.Request.Proto,
			params.StatusCode,
			params.Latency,
			params.Request.UserAgent(),
			params.ErrorMessage,
		)
	}))
	r.Use(gin.Recovery())
	r.GET("/ping",func(c *gin.Context){
		c.String(200,"pong")

	})
	r.Run()
}
