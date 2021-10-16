package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"net/http"
	"os"
)

func visit(links []string,n *html.Node)[]string{
	if n.Type == html.ElementNode && n.Data == "a"{
		for _,a := range n.Attr{
			if a.Key == "href"{
				links = append(links,a.Val)

			}
		}
	}

	for c := n.FirstChild;c != nil;c = c.NextSibling{
		links = visit(links,c)
	}
	return links

}

func main(){
	url := "http://www.baidu.com"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}
	//htmlcontent := fmt.Sprintf("%s", b)
	err = os.WriteFile("htmlfile",b,777)
	if err != nil{
		fmt.Println(err)
	}
	filename,err := os.Open("htmlfile")
	if err != nil{
		fmt.Println(err)
	}
	doc,err := html.Parse(filename)
	if err != nil{
		fmt.Fprintf(os.Stderr, "findlinks:%v\n",err)
		os.Exit(1)

	}
	for _,link := range visit(nil,doc){
		fmt.Println(link)
	}
}
