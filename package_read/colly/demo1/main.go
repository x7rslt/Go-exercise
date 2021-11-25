package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
)

func main() {
	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML(".pane-news", func(e *colly.HTMLElement) {
		fmt.Println(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("http://news.baidu.com/")
}
