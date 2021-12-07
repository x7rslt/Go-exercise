package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/post", func(c *gin.Context) {

		//id := c.Query("id")
		//page := c.DefaultQuery("page", "0")
		client := c.PostForm("client")
		linkman := c.PostForm("linkman")
		phone := c.PostForm("phone")
		email := c.PostForm("email")
		duration := c.PostForm("duration")
		types := c.PostForm("types")
		note := c.PostForm("note")

		fmt.Printf("The first data in monitor platform : 客户名称: %s; 联系人: %s; 手机号: %s; 邮箱: %s；服务期限：%s;服务类型:%s;备注:%s.\n", client, linkman, phone, email, duration, types, note)
	})
	router.Run(":8089")
}
