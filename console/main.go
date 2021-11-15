package main

import "fmt"

func main(){
	rst,err := ReadSingleLine()
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(rst)
	Write(rst)
	rstm,err := ReadPromptedLine("Read Prompt Test:")
	if err !=nil{
		fmt.Println(err)
	}
	fmt.Println(rstm)
	WriteMultipleLines("Write test from input:",[]string{"Hahaha","Lehehe"})

}