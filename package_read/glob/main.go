package main

import (
	"fmt"
	"github.com/gobwas/glob"
)

func main(){
	var g glob.Glob

	//Create simple glob
	g = glob.MustCompile("*.github.com")
	fmt.Println(g.Match("api.github.com"))

	g = glob.MustCompile(glob.QuoteMeta("*.github.com"))
	fmt.Println(g.Match("*.github.com"))

}