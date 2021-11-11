package main

import "fmt"

//fmt.Printf formatting tutorial

func main(){
	fmt.Printf("Binary: %b\\%b\n", 4, 5)   //Prints `Binary: 100\101`

	fmt.Printf("%d %%\n", 50) // Prints `50 %`

	//Use fmt.Sprintf to format a string without printing it,相当于赋值
	s := fmt.Sprintf("Binary: %b\\%b\n", 4, 5)  // s == `Binary: 100\101`
	fmt.Println(s)

	fmt.Printf("%q\n","Cafe")  //Prints Quoted string "Cafe"

	fmt.Printf("%6s\n","Cafe")  //Width 6, right justify

	fmt.Printf("%-6s\n","Cafe")  //Width 6, left justify

	fmt.Printf("%x\n","Cafe")   //Hex dump of byte values 43616665

	fmt.Printf("% x\n","Cafe")  //Hex dump with spaces 43 61 66 65




}