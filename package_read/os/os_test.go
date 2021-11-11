package os_test

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestOs(t *testing.T) {
	gp, err := os.Getgroups() //not supported by windows,see the os/user
	//package for a possible alternative
	fmt.Println(gp, err)

	pi := os.Getpagesize()
	fmt.Println(pi)

	cwd, err := os.Getwd()
	fmt.Println(cwd, err)

	hostname, err := os.Hostname()
	fmt.Println(hostname, err)

	filename := "os_test.go"
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Println("file does not exists.")
	}
	fmt.Println("file exists.")

	env, b := os.LookupEnv("path")
	fmt.Println(env, b)

	//os.MkdirAll("./test/mkdir/", 744) //create directory path

	homedir, err := os.UserHomeDir()
	fmt.Println(homedir, err)
}

func TestFile(t *testing.T) {
	f, err := os.OpenFile("newfile.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
	f, err = os.OpenFile("access.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write([]byte("appended some data\n")); err != nil {
		f.Close()
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

}

func TestDir(t *testing.T) {

	fmt.Println(os.Getwd())
	os.Chdir("..")
	fmt.Println(os.Getwd())
}
