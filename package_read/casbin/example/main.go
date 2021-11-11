package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"log"
	"os"
)

func main(){
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	modelfile := dir +"/package_read/casbin/example/model.conf"
	policyfile := dir + "/package_read/casbin/example/policy.csv"
	fmt.Println(modelfile,policyfile)
	e,err:= casbin.NewEnforcer(modelfile,policyfile)
	if err !=nil{
		log.Println("Casbin open err:",err)
	}
	sub := "alice" // the user that wants to access a resource.
	obj := "data1" // the resource that is going to be accessed.
	act := "read" // the operation that the user performs on the resource.

	if res, _ := e.Enforce(sub, obj, act); res {
		// permit alice to read data1
		log.Println("permit alice to read data1:",res)
	} else {
		// deny the request, show an error
		log.Println("deny the request, show an error:",res)
	}
	//今天身体不舒服，不知道为啥错，就想睡觉

}
