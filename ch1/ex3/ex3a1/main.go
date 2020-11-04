//练习1.3:尝试测量可能低效的程序和使用strings.Join的程序在执行时间上的差异。（1.6节有tmie包，11.4节展示如何撰写系统性的性能评估测试。）
package main

import ("fmt"
"os"
"time"
"strings"
)

func main(){
	s1 := time.Now()
	var s,sep string
	for i:= 1;i<len(os.Args);i++{
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	fmt.Printf("%.7fs done!",time.Since(s1).Seconds())
	fmt.Println("\t")
	s2 := time.Now()
	fmt.Println(strings.Join(os.Args[1:]," "))
	fmt.Printf("%.7fs done2!",time.Since(s2).Seconds())	
}