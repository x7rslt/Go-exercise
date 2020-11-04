//练习1.2:修改echo程序，输出参数的索引和值，每行一个
package main

import (
	"fmt"
	"os"
)

func main(){
	for num,arg := range os.Args[:]{
		fmt.Println(num,":",arg)
	}
}