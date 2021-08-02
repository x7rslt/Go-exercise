package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	router.GET("/someGet", MethodTest)
	router.POST("/somePost", MethodTest)
	router.PUT("/somePut", MethodTest)
	router.DELETE("/someDelete", MethodTest)
	router.PATCH("/somePatch", MethodTest)
	router.HEAD("/someHead", MethodTest)
	router.OPTIONS("/someOptions", MethodTest)

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run()
	// router.Run(":3000") for a hard coded port
}

func MethodTest(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"Message": c.Request.Header,
	})
}