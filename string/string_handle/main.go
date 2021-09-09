package main

import (
	"fmt"
	"math"
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
}