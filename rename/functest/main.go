package main

import (
	"fmt"
	"github.com/gookit/color"
	"os"
	"path/filepath"
	"strings"
)

func main(){
	//文件路径
	fmt.Println("请输入文件路径：")
	var filepath string
	fmt.Scanf("%s",&filepath)
	//文件名匹配特征
	match := "_"
	fmt.Println("文件列表如下：")
	filelist := FileList(filepath,match)
	var command string
	fmt.Println("是否重命名文件：Y/N ？")
	fmt.Scanf("%s",&command)
	if command !="y" {
		fmt.Println("取消重命名")
		os.Exit(0)
	}
	for _,name := range filelist{
		err := Rename(name,match)
		if err != nil{
			fmt.Println("重命名失败：",err)
		}
	}
	fmt.Println("All file rename done!")
}

//列出指定路径所有文件名
func FileList(basepath string,match string)[]string{
	var filelist []string
	err := filepath.Walk(basepath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		//匹配文件名规则，前缀"_"
		if !info.IsDir() && strings.HasPrefix(filepath.Base(path),match) {
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
func Rename(path string,match string)error{
	filename := filepath.Base(path)
	basepath:= filepath.Dir(path)
	//新命名规则
	newname := strings.TrimLeft(filename,match) //去除文件名前的"_"
	newpath := basepath + "/"+newname
	error := os.Rename(path,newpath)
	if error!= nil{
		return error
	}

	red := color.FgRed.Render
	green := color.FgGreen.Render
	fmt.Printf("原文件名：%s => 新文件名：%s\n",red(path),green(newpath))
	return nil
}