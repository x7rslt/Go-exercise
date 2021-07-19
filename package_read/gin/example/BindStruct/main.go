package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type StructA struct{
	Name string `form:"name"`

}
type StructB struct{
	NameStruct StructA
	Age int `form:"age"`

}

type StructC struct{
	NameStructPointer *StructA
	Sex string `form:"sex""`
}
type StructD struct{
	NameAnonyStruct struct{
		Love string `form:"love"`
	}
	Want string `form:"want"`
}

func GetDataA(c *gin.Context){
	var a StructA
	c.Bind(&a)
	c.JSON(http.StatusOK,gin.H{
		"a":a,
	})
}

func GetDataB(c *gin.Context){
	var b StructB
	c.Bind(&b)
	c.JSON(http.StatusOK,gin.H{
		"Name":b.NameStruct,
		"Age":b.Age,
	})
}

func GetDataC(c *gin.Context){
	var b StructC
	c.Bind(&b)
	c.JSON(http.StatusOK,gin.H{
		"a":b.NameStructPointer,
		"b":b.Sex,
	})
}
func GetDataD(c *gin.Context){
	var b StructD
	c.Bind(&b)
	c.JSON(http.StatusOK,gin.H{
		"a":b.NameAnonyStruct,
		"b":b.Want,
	})
}

func main(){
	r := gin.Default()
	r.GET("/geta",GetDataA)
	r.GET("/getb",GetDataB)
	r.GET("/getc",GetDataC)
	r.GET("/getd",GetDataD)
	r.Run()
}
// http://localhost:8080/geta?name=hello


