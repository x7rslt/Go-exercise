//修改dup2程序，输出出现重复行的文件的名称
package main
import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	list := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0{
		countLines("input",os.Stdin,list)
	}else{
		for _,arg := range files{
			f ,err := os.Open(arg)
			if err != nil{
				fmt.Fprintf(os.Stderr,"dup2:%v\n",err)
				continue
			}
			countLines(arg,f,list)
			f.Close()
		}
	}
	for filename,content := range list{
		for line,n := range content{
			if n > 1{
				fmt.Printf("%s has %d-lines: %s\n",filename,n,line)
			}
		}
	}
	}

func countLines(filename string,f *os.File,list map[string]map[string]int){
	edges := make(map[string]int)
	input := bufio.NewScanner(f)
	for input.Scan(){
		edges[input.Text()]++
	}
	list[filename] = edges
}