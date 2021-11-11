package ex3a2

import (
	"fmt"
	"os"
	"time"
	"strings"
)
//Echo1 is a low func
func Echo1(){
	s1 := time.Now()
	var s,sep string
	for i:= 1;i<len(os.Args);i++{
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	fmt.Printf("Echo1 %.7fs done!",time.Since(s1).Seconds())
}
//Echo2 is fast func
func Echo2(){	
	s2 := time.Now()
	fmt.Println(strings.Join(os.Args[1:]," "))
	fmt.Printf("Echo2 %.7fs done2!",time.Since(s2).Seconds())	
}