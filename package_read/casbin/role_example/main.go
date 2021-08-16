package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"log"
	"os"
)

func main(){
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	modelfile := path +"\\package_read\\casbin\\role_example\\model.conf"
	policyfile := path + "\\package_read\\casbin\\role_example\\policy.csv"
	fmt.Println(modelfile)
	e,err := casbin.NewEnforcer(modelfile,policyfile)
	if err != nil{
		log.Fatalf("Unable to create casbin enforcer:%v",err)

	}
	aliceRoles,err := e.GetRolesForUser("alice")
	if err !=nil{
		log.Fatalf("Could not get roles for alices:%v",err)

	}
	alicePerms,err := e.GetImplicitPermissionsForUser("alice")
	if err !=nil{
		log.Fatalf("Could not get permissions for alice:%v",err)

	}
	fmt.Printf(
		"alice is a member of the folloing roles:%v,and her permissions are :%v\n",
		aliceRoles,alicePerms,
		)
}