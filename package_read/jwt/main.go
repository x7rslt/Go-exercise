package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func main(){
	mySigningKey := []byte("FelixUtmostisthebest")

	type MyCustomClaims struct {
		Foo string `json:"foo"`
		jwt.StandardClaims
	}

	// Create the Claims
	claims := MyCustomClaims{
		"bar",
		jwt.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    "admin",
		},
	}
	//生成JWT-token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err!=nil{
		fmt.Println("JWT-token generate error:",err)
	}
	fmt.Printf("JWT-token : %v\n", ss)

	//解密JWT-token
	at(time.Unix(0, 0), func() {
		tokenString := ss
		token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte("FelixUtmostisthebest"), nil
		})

		if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
			fmt.Printf("%v %v", claims.Foo, claims.StandardClaims.Issuer)
		} else {
			fmt.Println(err)
		}
	})

}


// Override time value for tests.  Restore default value after.
func at(t time.Time, f func()) {
	jwt.TimeFunc = func() time.Time {
		return t
	}
	f()
	jwt.TimeFunc = time.Now
}