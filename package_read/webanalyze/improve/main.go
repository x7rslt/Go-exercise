package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	wappalyzer "github.com/projectdiscovery/wappalyzergo"
)

//响应快，结果没有原版"github.com/rverton/webanalyze"清晰
func main() {
	resp, err := http.DefaultClient.Get("https://www.freebuf.com/")
	if err != nil {
		log.Fatal(err)
	}
	data, _ := ioutil.ReadAll(resp.Body) // Ignoring error for example

	wappalyzerClient, err := wappalyzer.New()
	fingerprints := wappalyzerClient.Fingerprint(resp.Header, data)
	fmt.Printf("%v\n", fingerprints)

	// Output: map[Acquia Cloud Platform:{} Amazon EC2:{} Apache:{} Cloudflare:{} Drupal:{} PHP:{} Percona:{} React:{} Varnish:{}]
}

