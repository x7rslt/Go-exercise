package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var secrets = gin.H{
	"foo":gin.H{
		"email":"foo@bar.com","phone":"123456",
	},
	"austin":gin.H{"email":"austin@example.com","phone":"123456"},
	"lena":gin.H{"email":"lena@examlple.com","phone":"12345678"},
}

func main(){
	r := gin.Default()
	authorized := r.Group("/admin",gin.BasicAuth(gin.Accounts{
		"foo":"bar",
		"austin":"1234",
		"lena":"hello2",
		"manu":"4321",
	}))

	authorized.GET("/secrets",func(c *gin.Context){
		user:= c.MustGet(gin.AuthUserKey).(string)
		if secret,ok := secrets[user];ok{
			c.JSON(http.StatusOK,gin.H{
				"user":user,"secrets":secret,
			})
		}else{
			c.JSON(http.StatusOK,gin.H{
				"user":user,"secret":"No secret:(",
			})
		}
	})

	r.Run()

}