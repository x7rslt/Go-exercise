package main

import (
	"bufio"
	"crypto/tls"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

type JsonReturn struct {
	Url     string
	Pattern string
	Match   string
}
//待处理
func getLeak(url string, data string, pattern string, jsonArray *[]JsonReturn) {
	re := regexp.MustCompile(`^[a-z]+\[[0-9]+\]$`)
	matches := re.FindStringIndex(data)
	//fmt.Println(len(matches))
	if len(matches) != 0 {
		fmt.Printf("[+] Url: %v\n[+] Pattern: %v\n[+] Match: %v\n", url, pattern, string(len(matches)))
		jsn := JsonReturn{url, pattern, string(len(matches))}
		*jsonArray = append(*jsonArray, jsn)
	}

	/*
	re := regexp.MustCompile(`^[a-z]+\[[0-9]+\]$`) //regex equation

		fmt.Println(re.MatchString("adam[23]")) //判断是否匹配
	 */
}

func get_inputs() []string {
	reader := bufio.NewReader(os.Stdin)
	var output []rune

	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		output = append(output, input)
	}

	return strings.Fields(string(output))
}

func req(url string, timeout int) string {
	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
	}
	client := &http.Client{
		Transport: transCfg,
		Timeout: time.Duration(timeout) * time.Second,
	}
	res, err := client.Get(url)

	if err != nil {
		log.Fatal(err)
	}
	data, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	return string(data)
}

func main() {
	path := flag.String("pattern", "", "[+] File contains patterns to test")
	verbose := flag.Bool("verbose", false, "[+] Verbose Mode")
	jsonOutput := flag.String("json", "", "[+] Json output file")
	timeout := flag.Int("timeout", 5, "[+] Timeout for request in seconds")
	flag.Parse()

	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		fmt.Println("[+] Use in Pipeline")
		os.Exit(1)
	}

	file, err := os.Open(*path)
	defer file.Close()
	lines := make([]string, 0)

	patterns := bufio.NewScanner(file)
	jsonArray := make([]JsonReturn, 1)
	for patterns.Scan() {
		lines = append(lines, patterns.Text())
	}

	if err != nil {
		fmt.Println("Open error:",err)
		//log.Fatal(err)
	}

	for _, url := range get_inputs() {
		if *verbose {
			fmt.Println("[-] Looking: " + url)
		}

		data := req(url,*timeout)

		for _, pattern := range lines {
			getLeak(url, data, pattern, &jsonArray)
		}
	}

	if *jsonOutput != "" {
		fo, err2 := os.Create(*jsonOutput)
		k, err1 := json.MarshalIndent(jsonArray, "", "\t")
		if _, err := fo.Write(k); err1 != nil || err2 != nil {
			panic(err)
		}
	}
}
