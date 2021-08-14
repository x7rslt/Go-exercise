package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)


type MyStandardClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func main(){
	//MapClaims map
	//StandardClaims struct
	mySigningKey := []byte("Helloworld")
	c:= MyStandardClaims{
		Username:"x7rslt",
		StandardClaims:jwt.StandardClaims{
			NotBefore: time.Now().Unix()-60,
			ExpiresAt: time.Now().Unix()+60*60*2,
			Issuer:"admin",
		},
	}
	t :=jwt.NewWithClaims(jwt.SigningMethodHS256,c)
	jwttoken,err := t.SignedString(mySigningKey)
	if err != nil{
		fmt.Println("JWT token generate error:",err)
	}
	fmt.Println(jwttoken)

	token,err := jwt.ParseWithClaims(jwttoken,&MyStandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey,nil
	})
	if err != nil{
		fmt.Println("JWT token decode error:",err)
		return
	}
	fmt.Println(token.Claims.(*MyStandardClaims).Username)

}