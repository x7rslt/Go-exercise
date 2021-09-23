package type_test

import (
	"fmt"
	"reflect"
	"testing"
)


//Use fmt for a string type description
func TestFmtType(t *testing.T){
	var x interface{} = []int{1, 2, 3}
	xType := fmt.Sprintf("%T", x)
	fmt.Println(xType) // "[]int"
}


//A type switch lets you choose between types
func TestSwitchType(t *testing.T){
	var x interface{} = 2.3
	switch v := x.(type) {
	case int:
		fmt.Println("int:", v)
	case float64:
		fmt.Println("float64:", v)
	default:
		fmt.Println("unknown")
	}
	// Output: float64: 2.3
}

//Reflection gives full type information
func TestReflection(t *testing.T){
	var x interface{} = []int{1, 2, 3}
	xType := reflect.TypeOf(x)
	xValue := reflect.ValueOf(x)
	fmt.Println(xType, xValue) // "[]int [1 2 3]"
}