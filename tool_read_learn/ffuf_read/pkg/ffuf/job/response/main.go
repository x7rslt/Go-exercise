package main

import "net/http"

type Request struct {
	Method string
	Host string
	Url string
	Headers map[string]string
	Data []byte
	Input map[string][]byte
	Position int
	Raw string
}

func NewRequest()Request{
	var req Request
	req.Method = "Get"
	req.Url = "www.baidu.com"
	req.Headers = Make(map[string]string)
	return req
}

type Response struct{
	StatusCode int64
	Headers map[string[]]string
	Data []byte
	Contentlength int64
	ContentWords int64
	ContentLines int64
	Cancelled bool
	Request *Request
	Raw string
	ResultFile string
}

func NewResponse(httpresp *http.Response, req *Request) Response {
	var resp Response
	resp.Request = req
	resp.StatusCode = int64(httpresp.StatusCode)
	resp.Headers = httpresp.Header
	resp.Cancelled = false
	resp.Raw = ""
	resp.ResultFile = ""
	return resp
}

func main() {
	var httpreq *http.Request
	data := bytes.NewReader([]byte("test"))
	q := NewRequest()
	fmt.Println(q)

}
