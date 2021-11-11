package main

import "fmt"

type vector[T any] []T

func printSlice[T any](s []T){
	for _,v := range s{
		fmt.Printf("%v ",v)

	}
	fmt.Println("\n")
}

func main(){
	v := vector[int]{58,188}
	printSlice(v)
	v2:=vector[string]{"红烧肉","九转大肠"}
	printSlice(v2)
}