package main

import "fmt"

type Addable interface {
	type int,string
}

func add[T Addable](a,b T)T{
	return a +b
}

func main(){
	fmt.Println(add(3,4))
	fmt.Println(add("Hellow","world!"))
}
