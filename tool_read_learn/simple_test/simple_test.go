package simple

import (
	"fmt"
	"strings"
	"testing"
)

//main.go test wordlist Set()
func Test(t *testing.T) {
	fmt.Println("lehehe")
	value := "simple,test"
	delimited := strings.Split(value, ",")
	fmt.Println(delimited)
}
