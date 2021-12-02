package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
)
//No license found, subdomain scan function can not be used
var (
	file      = flag.String("f", "conf/domain.txt", "域名文件")
	thread   = flag.Int("t", 3, "进程数 例如:-n=3")
	h         = flag.Bool("h", false, "帮助信息")

)

func main() {
	flag.Parse()
	if *h == true {
		flag.PrintDefaults()
		return
	}
	domain_file, err := os.Open(*file)
	if err != nil{
		fmt.Println(err)
	}
	defer domain_file.Close()
	sc := bufio.NewScanner(domain_file)
	var domains []string
	for sc.Scan() {
		domain := sc.Text()
		domains = append(domains, domain)
		//fmt.Println(domains)
	}
	c_task := make(chan struct{}, len(domains))
	c_thread := make(chan struct{}, *thread)
	defer close(c_task)
	defer close(c_thread)
	for _, domain := range domains{
		fmt.Println(domain)
		c_thread <- struct{}{}
		c_task <- struct{}{}
		go func(domain string) {
			fmt.Println("Scan:",domain)
			cmd := exec.Command("./xray_darwin_arm64", "subdomain", "--target", domain, "--json-output", domain+".json")
			output, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Println(fmt.Sprint(err) + ":" + string(output))
				return
			}
			<- c_thread

		}(domain)
	}
	for i:=len(domains); i>0;i-- {
		<- c_task
	}
	//time.Sleep(time.Second)  待解决：多线程无等待，goroutine没有执行
	fmt.Println("主协程")

}
