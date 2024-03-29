package string_test

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"testing"
)

func Test(t *testing.T) {
	str := "Go 爱好者"
	log.Printf("The string:%q\n", str)
	fmt.Printf("=>runes(char):%q\n", []rune(str))
	fmt.Printf("=>runes(hex):%x\n", []rune(str))
	fmt.Printf("=>bytes(hex):[% x]\n", []byte(str))
	fmt.Printf("str 有%d个字节\n", len(str))
	fmt.Printf("str 有%d个字符\n", len([]rune(str)))
	for i, j := range str {
		fmt.Println(i, j)
		fmt.Printf("%d,%q\n", i, j)
	}
}

//定义IP地址的打印格式
type IP [4]int

func (i IP) String() {
	fmt.Sprintf("IP:%d.%d.%d.%d", i[0], i[1], i[2], i[3])
}

func TestIP(t *testing.T) {
	localhost := IP{127, 0, 0, 1}
	localhost.String()
	hosts := map[string]IP{
		"localhost": {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v:%v\n", name, ip)
	}

}


func TestPasswordList(t *testing.T){
	path,_ := os.Getwd()
	file := path + "/password.txt"
	f,err := os.Open(file)
	if err != nil{
		fmt.Println("Open file error:",err)
	}
	defer f.Close()
	var passwordlist []string
	//逐行读取string
	scanner := bufio.NewScanner(f)
	for scanner.Scan(){
		passwordlist = append(passwordlist, "\""+scanner.Text()+"\""+",")
	}
	if err := scanner.Err();err!= nil{
		fmt.Println("Read file err:",err)
	}
	fmt.Println(passwordlist)
	//逐行写入string
	newfile,err := os.Create(path+"/newlist.txt")
	if err != nil{
		fmt.Println("Create File err:",err)

	}

	defer newfile.Close()

	for _,line := range passwordlist{
		newfile.WriteString(line)
	}
}

//Alternate password
func TestPasswdGenerate(t *testing.T){
	path,_ := os.Getwd()
	file := path + "/password.txt"
	f,err := os.Open(file)
	if err != nil{
		fmt.Println("Open file error:",err)
	}
	defer f.Close()
	var passwordlist []string
	//逐行读取string
	scanner := bufio.NewScanner(f)
	for scanner.Scan(){
		passwordlist = append(passwordlist, scanner.Text()+"\n")
	}
	if err := scanner.Err();err!= nil{
		fmt.Println("Read file err:",err)
	}
	fmt.Println(passwordlist)
	//逐行写入string
	newfile,err := os.Create(path+"/alterlist.txt")
	if err != nil{
		fmt.Println("Create File err:",err)

	}

	defer newfile.Close()

	for i,line := range passwordlist{
		println(i)
		newfile.WriteString(line)
		newfile.WriteString("peter\n")
	}
}