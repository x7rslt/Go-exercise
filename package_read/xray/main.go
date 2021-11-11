package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"
	"xrayscan/common"
)
type XrayResAll struct {
	CreateTime int64 `json:"create_time"`
	Detail     struct {
		Addr     string     `json:"addr"`
		Payload  string     `json:"payload"`
		Snapshot [][]string `json:"snapshot"`
		Extra    struct {
			Author string `json:"author"`
			Param  struct {
			} `json:"param"`
		} `json:"extra"`
	} `json:"detail"`
	Plugin string `json:"plugin"`
	Target struct {
		URL string `json:"url"`
	} `json:"target"`
}
var (
	lock sync.Mutex
)

var targets []string = []string{
	"http://testphp.vulnweb.com/",
}

func main(){
	err:= ScanPocXray("",targets)
	if err!=nil{
		log.Printf("Scan error:%v",err)
	}
}
func ScanPocXray(pocName string, targetString []string) error {

	osType := runtime.GOOS
	nowPath, _ := os.Getwd()

	if _, err := os.Stat("output"); os.IsNotExist(err) {
		_ = os.Mkdir("output", os.ModePerm)
	}
	if _, err := os.Stat(filepath.Join("output", "xrayinput")); os.IsNotExist(err) {
		_ = os.Mkdir(filepath.Join("output", "xrayinput"), os.ModePerm)
	}
	if _, err := os.Stat(filepath.Join("output", "xrayout")); os.IsNotExist(err) {
		_ = os.Mkdir(filepath.Join("output", "xrayout"), os.ModePerm)
	}

	xrayOuputPaht := filepath.Join(nowPath, "output", "xrayinput")
	fileName := common.GetRandomString(12)

	fileName = filepath.Join(xrayOuputPaht, fileName+".txt")
	common.WirteFileAppend(fileName, targetString)

	// 输出文件路径
	outJson := time.Now().Format("2006-01-02 15:04:05") + ".json"
	xrayOuputputh := filepath.Join(nowPath, "output", "xrayout", outJson)

	cmd := &exec.Cmd{}

	if osType == "darwin" {

		// 要运行的程序名称，比如 ksubdomainMac
		runName := filepath.Join(nowPath, "third", "xray_darwin_amd64")
		log.Println("Run path:",runName)
		//./xray_darwin_amd64 webscan --url http://127.0.0.1:19001  --html-output 0102.html
		// --webhook-output http://127.0.0.1:5000/webhook
		if pocName != "" {
			cmd = exec.Command(runName, "webscan", "--url-file", fileName, "--poc", pocName, "--json-output", xrayOuputputh)
		} else {
			cmd = exec.Command(runName, "webscan", "--url-file", fileName, "--json-output", xrayOuputputh)
		}

	} else if osType == "linux" {
		runName := filepath.Join(nowPath, "pkg", "third", "xray")

		if pocName != "" {
			cmd = exec.Command(runName, "webscan", "--url-file", fileName, "--poc", pocName, "--json-output", xrayOuputputh)
		} else {
			cmd = exec.Command(runName, "webscan", "--url-file", fileName, "--json-output", xrayOuputputh)
		}
	} else {
		runName := filepath.Join(nowPath, "third", "xray")
		log.Println("Runname:",runName)
		//./xray_darwin_amd64 webscan --url http://127.0.0.1:19001  --html-output 0102.html
		// --webhook-output http://127.0.0.1:5000/webhook
		if pocName != "" {
			cmd = exec.Command(runName, "webscan", "--url-file", fileName, "--poc", pocName, "--json-output", xrayOuputputh)
		} else {
			cmd = exec.Command(runName, "webscan", "--url-file", fileName, "--json-output", xrayOuputputh)
		}
	}

	fmt.Println("\n\nxray", cmd)

	cmdOutput, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("runcmd err : " + fmt.Sprint(err) + ":" + string(cmdOutput))
	}

	//if common.PathExist(fileName) {
	//	os.Remove(fileName)
	//}
	//
	if common.PathExist(xrayOuputputh) {
		XrayRes(xrayOuputputh)
		//os.Remove(xrayOuputputh)
	}

	return nil

}

func XrayRes(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		return
	}
	defer f.Close()
	buf := bufio.NewReader(f)
	for {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("xray into db err :", err)
			}
		}()
		var xrayResult XrayResAll

		lines, err := buf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				return
			}
			return
		}
		decoder := json.NewDecoder(strings.NewReader(lines))
		err = decoder.Decode(&xrayResult)
		if err != nil {
			fmt.Println("Decoder failed : ", err.Error())
		} else {
			snapshot := common.SliceSToString(xrayResult.Detail.Snapshot)
			hash := common.GetMD5Hash(xrayResult.Target.URL+xrayResult.Plugin)
			lock.Lock()
			data := make(map[string]interface{})
			data["url"] = xrayResult.Target.URL
			data["poc"] = xrayResult.Plugin
			data["hash"] = hash
			data["snapshot"] = snapshot
			log.Print(data)
			lock.Unlock()
		}
	}
}
