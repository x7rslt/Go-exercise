package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	filepath2 "path/filepath"
	"strings"
	"unicode/utf8"
)

func main(){
	fmt.Println(" ")
	fmt.Println("Japan" == "Japan")
	fmt.Println(strings.EqualFold("Japan", "JAPAN"))

	//Length in bytes or runes
	fmt.Println(len("日"))
	fmt.Println(utf8.RuneCountInString("日"))
	fmt.Println(utf8.ValidString("日"))
	fmt.Sprintf("%.4f", math.Pi)


	abspath,_ := filepath2.Abs("./string/string_handle")
	file,err := os.Open(abspath+"/content")
	if err != nil{
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		fmt.Println(scanner.Text())
		fmt.Println(strings.TrimSpace(scanner.Text()))
	}


}