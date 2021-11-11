package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Declare http client
	client := &http.Client{}

	// Declare post data
	//PostData := strings.NewReader("useId=5&age=12")

	// Declare HTTP Method and Url
	req, err := http.NewRequest("GET", "https://www.gezhong.vip/tree.php", nil)

	// Set cookie
	req.Header.Set("Cookie", "__cfduid=d0dcf8f93f615e47c81f92c2a9ca5add11607150369; __gads=ID=d7e02b7bc6ed9d5e-22ddd84c06c500f7:T=1607150371:RT=1607150371:S=ALNI_MbQJ2wOEqytPVOe540rPxH_guY0HA; bbs_sid=t4echkd6lp3vel673vh3ha5th2; Hm_lvt_7dec70fe7ef7916dc22570b81f0d96bd=1607150372,1607150461,1608816728; logined=1; bbs_token=W0ib3KaNkq8k9GTSqkY8wMWIM9V_2Bk5F0JkbZbT5bDgDVw9aIKktQ2gMimC2KYzeBh7FaQsLtZQEyeWmjBgow8FdFt8U_3D; Hm_lpvt_7dec70fe7ef7916dc22570b81f0d96bd=1608816820")
	resp, err := client.Do(req)
	// Read response
	data, err := ioutil.ReadAll(resp.Body)

	// error handle
	if err != nil {
		fmt.Printf("error = %s \n", err)
	}

	// Print response
	fmt.Printf("Response = %s", string(data))
}
