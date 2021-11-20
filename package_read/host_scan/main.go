package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"github.com/fatih/color"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type info struct {
	ip string
	host string
	url string
	status int
	length int
	title string
	location string
	content []byte
}

func ScreenPrint(i info){
	cyan := color.New(color.FgBlue).PrintfFunc()
	green := color.New(color.FgGreen).PrintfFunc()
	yellow := color.New(color.FgYellow).PrintfFunc()
	red := color.New(color.FgRed).PrintfFunc()
	var status = strconv.Itoa(i.status)
	title := i.title
	location := i.location
	length := strconv.Itoa(i.length)
	if status[0] == '2'{
		green("Status: %s ",status)
		green("Title: %s ",title)
		//green("Location: %s ",location)
		green("Length:%s\n ",length)
	}else if status[0] == '3'{
		yellow("Status: %s ",status)
		yellow("Title: %s ",title)
		yellow("Location: %s ",location)
	}else{
		red("Status: %s ",status)
		red("Title: %s ",title)
		red("Length:%s ",length)
	}

	if i.status/100 == 3{
		cyan("%s\t%s\t%s\t%s\t%s\t%s\n",i.ip,i.url,status, length, title,location)
	}else{
		cyan("Info:%s\t%s\t%s\t%s\t%s\n",i.ip,i.url,status,length,title)
	}
}

func sendReq(client *http.Client,ip string,host string,path string)(ret []info){
	schemaHttp := "http://"
	schemaHttps := "https://"
	for _,schema:= range[]string{schemaHttp,schemaHttps}{
		req,_:=http.NewRequest(http.MethodGet,schema+ip+path,nil)
		req.Host = host
		req.Header.Set("User-Agent",userAgent)
		resp,err:= client.Do(req)
		if err!= nil{
			continue
		}
		content,_ := io.ReadAll(resp.Body)
		var location = ""
		if resp.StatusCode/100 == 3{
			location = resp.Header.Get("Location")

		}
		ret = append(ret,info{ip,host,schema+host+path,resp.StatusCode,len(content),getTitle(content),location, content})
	}
	return ret
}
func getTitle(content []byte)(title string){
	reg,_:=regexp.Compile(`(?Ui:<title>[\s ]*([\s\S]*)[\s ]*</?title>)`)
	m:= reg.FindStringSubmatch(string(content))
	if len(m)!=0{
		title = strings.Replace(m[1], "\n", "", -1)
	}
	return title
}
var userAgent string

func FileExist(path string)bool{
	_,err := os.Lstat(path)
	return !os.IsNotExist(err)
}

func writeFile(w *bufio.Writer,i info){
	var line string
	if i.status/100 == 3{
		line =fmt.Sprintf("%s: %s\t%s: %s\t%s: %d\t%s: %d\t%s: %s\t%s: %s\n",
			"IP",i.ip,
			"URL",i.url,
			"Status",i.status,
			"Length",i.length,
			"Title",i.title,
			//cyan("Type"),i.Type,
			"Location",i.location,
		)
	}else{
		line = fmt.Sprintf("%s: %s\t%s: %s\t%s: %d\t%s: %d\t%s: %s\n",
			"IP",i.ip,
			"URL",i.url,
			"Status",i.status,
			"Length",i.length,
			"Title",i.title,
			//cyan("Type"),i.Type,
		)
	}
	_,err := w.WriteString(line)
	if err!=nil{
		panic(err)
	}
	_= w.Flush()
}

func main(){
	userAgent = "This is a test."
	var timeout int
	baiduip := info{"127.0.0.1","10.0.0.1","http://www.baidu.com",200,400,"百度科技","Beijing",[]byte{}}
	ScreenPrint(baiduip)
	//skip verify cert
	tr := &http.Transport{
		TLSClientConfig:&tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Timeout: time.Duration(timeout)*time.Second,
		Transport:tr,
		CheckRedirect: func(req *http.Request,via []*http.Request)error{
			return http.ErrUseLastResponse
		},
	}
	result:= sendReq(client,"180.101.49.11","180.101.49.11","/")
	fmt.Println(result[0].title)
	path,_ := filepath.Abs("./package_read/host_scan")
	f, err := os.OpenFile(path+"/ScanResult", os.O_CREATE|os.O_WRONLY, 0664)
	if err!=nil{
		panic(err)
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	writeFile(w,result[0])


}