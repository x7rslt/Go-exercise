package main

import (
	"fmt"
)

import "github.com/bobesa/go-domain-util/domainutil"

func main(){
	fmt.Println(domainutil.Domain("keep.google.com"))
	fmt.Println(domainutil.DomainSuffix("news.baidu.com"))
	fmt.Println(domainutil.Subdomain("wechat.tencetn.com"))
	fmt.Println(domainutil.HasSubdomain("test.baidu.com"))
}
