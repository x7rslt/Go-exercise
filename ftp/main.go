package main

import "net/http"

func main() {
	h := http.FileServer(http.Dir("/")) //开放目录
	http.ListenAndServe(":8888", h)
}
