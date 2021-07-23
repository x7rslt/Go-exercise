package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path/filepath"
)

func main(){
	r := gin.Default()
	r.MaxMultipartMemory=8<<20
	r.Static("/","./public")
	r.POST("/upload",func(c *gin.Context){
		name := c.PostForm("name")
		email := c.PostForm("email")
		file,_:= c.FormFile("file")
		log.Println(file.Filename)
		filepath := "./uploadfile/"+filepath.Base(file.Filename)
		log.Println(filepath)
		c.SaveUploadedFile(file,filepath)
		c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully with fields name=%s and email=%s.", file.Filename, name, email))

	})
	r.Run()

}