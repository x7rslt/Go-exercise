package rename_test

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRun(t *testing.T){

	filepath := "/Users/felix/CodeDev/GoDatabase/Go-exercise/rename/filetest"
	//Getfilelist(filepath)
	fmt.Println("文件列表如下：")
	filelist := FileList(filepath)
	var command string
	fmt.Println("是否重命名文件：Y/N ？")
	fmt.Scanf("%s",&command)
	if command !="y" {
		fmt.Println("取消重命名")
		os.Exit(0)
	}
	for _,name := range filelist{
		err := Rename(name)
		if err != nil{
			fmt.Println("重命名失败：",err)
		}
	}
	fmt.Println("All file rename done!")
}

//列出指定路径所有文件名
func FileList(basepath string)[]string{
	var filelist []string
	err := filepath.Walk(basepath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		//匹配文件名规则，前缀"_"
		if !info.IsDir() && strings.HasPrefix(filepath.Base(path),"_") {
			fmt.Printf("文件名: %s\n",path)
			filelist = append(filelist, path)
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
	return filelist
}

//重命名
func Rename(path string)error{
	filename := filepath.Base(path)
	basepath:= filepath.Dir(path)
	//新命名规则
	newname := strings.TrimLeft(filename,"_") //去除文件名前的"_"
	newpath := basepath + "/"+newname
	error := os.Rename(path,newpath)
	if error!= nil{
		return error
	}
	return nil
}