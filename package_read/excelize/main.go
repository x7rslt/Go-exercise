package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"log"
	"os"
)

func main(){
	ReadExcel("ExcelWrite2.xlsx")

}
func ReadExcel(filename string){
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	path = path + "/package_read/excelize/"+filename
	f2,err := excelize.OpenFile(path)
	if err != nil{
		fmt.Println(err)
		return
	}
	//获取工作表中指定单元格的值
	cell,err := f2.GetCellValue("Sheet1","B2")
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(cell)

	//获取Sheet1 上所有单元格
	rows,err := f2.GetRows("Sheet1")
	if err != nil{
		fmt.Println(err)
		return
	}
	//批量打印所有值
	fmt.Println(rows)
	//逐行打印
	for line,row := range rows{
		fmt.Printf("Line %v,Value %v\n",line,row)
		for _,colCell := range row{
			fmt.Println(colCell,"\t")
		}
		fmt.Println()
	}


}

func WriteExcel(){
	f := excelize.NewFile()
	//Create a new sheet.
	index := f.NewSheet("Sheet2")
	//Set value of a cell
	f.SetCellValue("Sheet2","A2","Hello world.")
	f.SetCellValue("Sheet1","B2",100)
	f.SetActiveSheet(index)
	//Current path
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Current path :",path)
	if err := f.SaveAs(path + "/package_read/excelize/ExcelWrite2.xlsx");err != nil{
		fmt.Println(err)
	}
}