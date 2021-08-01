package main

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"time"
)

var (g errgroup.Group)

func R1()http.Handler{
	e:=gin.New()
	e.Use(gin.Recovery())
	e.GET("/",func(c *gin.Context){
		c.JSON(
			http.StatusOK,
			gin.H{
				"code":http.StatusFound,"message":"Welcome server 01",
			})
	})
	return e
}

func R2()http.Handler{
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/",func(c *gin.Context){
		c.JSON(
			http.StatusOK,
			gin.H{
				"code":http.StatusOK,
				"message":"Welcome server 02",

			})
	})
	return e
}

func main(){
	server01 := &http.Server{
		Addr:":8080",
		Handler: R1(),
		ReadTimeout: 5*time.Second,
		WriteTimeout: 10*time.Second,
	}

	server02 := &http.Server{
		Addr:":8081",
		Handler: R2(),
		ReadTimeout: 5*time.Second,
		WriteTimeout: 10*time.Second,
	}
	g.Go(func() error {
		return server01.ListenAndServe()
	})
	g.Go(func()error{
		return server02.ListenAndServe()
	})
	if err := g.Wait();err!=nil{
		log.Fatal(err)
	}
}
