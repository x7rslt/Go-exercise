package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)




func main(){
	//const timeout = time.Duration(1 * 1000000)

	var tr = &http.Transport{
		MaxIdleConns:      30,
		IdleConnTimeout:   time.Second,
		DisableKeepAlives: true,
		TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
		DialContext: (&net.Dialer{
			//Timeout:   timeout,
			KeepAlive: time.Second,
		}).DialContext,
	}
	re := func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}
	client := &http.Client{
		Transport:     tr,
		CheckRedirect: re,
		//Timeout:       timeout,
	}

	url := "https://www.baidu.com"
	result := isListening(client,url,"GET")
	fmt.Println(result)

}


func isListening(client *http.Client, url, method string) bool {


	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return false
	}

	req.Header.Add("Connection", "close")
	req.Close = true

	resp, err := client.Do(req)
	if resp != nil {
		fmt.Println(resp.Header)
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
	}

	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

