package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	url := "https://drive.google.com/file/d/1G1t4MvLPJARf2PHP53npo4clMZfHk1wn"
	pwd, _ := os.Getwd()
	resp, _ := http.Get(url)

	fmt.Println(url, pwd, resp)
}
