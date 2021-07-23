package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

func main(){
	r := gin.Default()
	r.MaxMultipartMemory = 8<<20
	r.Static("/","./public")
	r.POST("/upload",func(c *gin.Context){
		name := c.PostForm("name")
		email := c.PostForm("email")

		form,err := c.MultipartForm()
		if err != nil{
			c.String(http.StatusBadRequest,fmt.Sprintf("get form err:%s",err.Error()))
			return
		}
		files := form.File["files"]
		for _,file := range files{
			filename := filepath.Base(file.Filename)
			if err := c.SaveUploadedFile(file,filename);err!=nil{
				c.String(http.StatusBadRequest,fmt.Sprintf("upload file err:%s",err.Error()))
				return
			}
		}
		c.String(http.StatusOK,fmt.Sprintf("Upload success %d files with fields name= %s and email=%s.",len(files),name,email))

	})
	r.Run()
}