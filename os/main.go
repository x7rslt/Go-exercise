package main

import (
	"fmt"
	"os/user"
)

func main(){
	username,err := user.Current()
	if err != nil{
		fmt.Println(err)
	}

	fmt.Println("Current user name is : ",username.Username)

	isuser,_ := user.Lookup("test")
	fmt.Println("User group has user 'test'",isuser)
}


