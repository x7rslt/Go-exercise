package os_test

import (
	"fmt"
	"log"
	"os/exec"
	"testing"
)

func TestStatStdin(t *testing.T) {
	url := "http://www.baidu.com"
	fmt.Println(OpenUrl(url))


}

func OpenUrl(url string) []byte{
	out, err := exec.Command("open", url).Output()
	if err != nil {
		log.Fatal(err)
	}
	return out

}