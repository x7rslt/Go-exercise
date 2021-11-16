package main

import (
	"github.com/fatih/color"
	"github.com/rodaine/table"
)

func main(){
	headerFmt := color.New(color.FgGreen,color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("ID","Name","Score","Added")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	getWidgets := map[string]string{
		"ID":"10010",
		"Name":"Hahaha",
		"Cost":"10000",
		"Added":"100000",
	}
	tbl.AddRow(getWidgets["ID"],getWidgets["Name"],getWidgets["Cost"],getWidgets["Added"])
	tbl.Print()
}