package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
)

var tplString=`
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>{{.Title}}</title>
</head>
<body>
</body>
</html>
`

type Data struct {
	Title    string
}

func checkErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

func ParseString(data Data)  {
	var err error
	var t *template.Template
	t = template.New("Products") //创建一个模板
	t, err = t.Parse(tplString)
	checkErr(err)
	err = t.Execute(os.Stdout, data)
	checkErr(err)
}
func ParseFileWrong(data Data)  {
	var err error
	var t *template.Template
	t = template.New("Products") //创建一个模板
	t, err = t.ParseFiles("tpl.html")
	checkErr(err)
	err = t.Execute(os.Stdout, data)
	checkErr(err)
}
func ParseFile(data Data)  {
	var err error
	var t *template.Template
	t, err = template.ParseFiles("tpl.html")  //从文件创建一个模板
	checkErr(err)
	err = t.Execute(os.Stdout, data)
	checkErr(err)
}

func main() {
	data:=Data{Title:"夕阳西下"}
	fmt.Println("-----第一种----")
	ParseString(data)
	fmt.Println("-----第二种（错误）----")
	ParseFileWrong(data)
	fmt.Println("-----第三种----")
	ParseFile(data)
}