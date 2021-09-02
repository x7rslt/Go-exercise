package path_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func getExecutePath1() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(dir)

	return dir
}

func getExecutePath2() string {
	dir, err := os.Executable()
	if err != nil {
		fmt.Println(err)
	}

	exPath := filepath.Dir(dir)
	//fmt.Println(exPath)

	return exPath
}

func getExecutePath3() string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(dir)

	return dir
}

func getExecutePath4() string {
	dir, err := filepath.Abs("./")
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(dir)
	return dir
}

func TestPath(t *testing.T){

	//解释器执行的路径
	path1 := getExecutePath1()
	fmt.Println("Path1:",path1)
	//同上
	path2 := getExecutePath2()
	fmt.Println("Path2:",path2)
	//main文件的路径
	path3 := getExecutePath3()
	fmt.Println("Path3:",path3)
	//同上
	path4 := getExecutePath4()
	fmt.Println("Path4:",path4)

}