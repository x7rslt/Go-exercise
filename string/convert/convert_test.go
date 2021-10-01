package convert_test

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"testing"
)


//Split string to slice
func TestString2Slice(t *testing.T){
	//Method 1
	s := strings.Split("a,b,c",",")
	fmt.Println("s:",s)

	//Method 2: split by whitespace and newline
	s2:= strings.Fields(" a\t b \n")
	fmt.Println("S2:",s2)

	//Method3: split on regular expression
	a := regexp.MustCompile(`a`)              // a single `a`
	fmt.Printf("%q\n", a.Split("banana", -1)) // ["b" "n" "n" ""]
	fmt.Printf("%q\n", a.Split("banana", 0))  // [] (nil slice)
	fmt.Printf("%q\n", a.Split("banana", 1))  // ["banana"]
	fmt.Printf("%q\n", a.Split("banana", 2))  // ["b" "nana"]

	zp := regexp.MustCompile(` *, *`)             // spaces and one comma
	fmt.Printf("%q\n", zp.Split("a,b ,  c ", -1)) // ["a" "b" "c "]

}

//Convert string and bytes
func TestStringByte(t *testing.T){
	//Convert string to bytes
	b := []byte("ABC€")
	fmt.Println(b) // [65 66 67 226 130 172]

	//Convert bytes to string
	s := string([]byte{65, 66, 67, 226, 130, 172})
	fmt.Println(s) // ABC€

	h := []byte("Hello world")
	fmt.Println(h)

	s2 := string(h)
	fmt.Println(s2)
}


//Convert between rune array/slice and string
func TestStringRune(t *testing.T){
	//Convert string to runes
	r := []rune("ABC€")
	fmt.Println(r)        // [65 66 67 8364]
	fmt.Printf("%U\n", r) // [U+0041 U+0042 U+0043 U+20AC]

	//Convert runes to string
	s := string([]rune{'\u0041', '\u0042', '\u0043', '\u20AC', -2})
	fmt.Println(s) // ABC€�
}

//Convert between float and string
func TestStringFloat(t *testing.T){
	//String to float
	f := "3.14159265"
	if s, err := strconv.ParseFloat(f, 32); err == nil {
		fmt.Println(s) // 3.1415927410125732
	}
	if s, err := strconv.ParseFloat(f, 64); err == nil {
		fmt.Println(s) // 3.14159265
	}

	//Float to string
	s:= fmt.Sprintf("%f",123.456) // s == "123.456000"
	fmt.Println(s)
}

//Convert between int, int64 and string
func TestStringInt(t *testing.T){
	//int/int64 to string
	s := strconv.Itoa(97) // s == "97"
	fmt.Println(s)

	//Int64
	var n3 int64 = 97
	s3 := strconv.FormatInt(n3, 10) // s == "97" (decimal)
	fmt.Println(s3)

	var n4 int64 = 97
	s4 := strconv.FormatInt(n4, 16) // s == "61" (hexadecimal)
	fmt.Println(s4)

	//string to int/int64
	s5 := "97"
	if n, err := strconv.Atoi(s5); err == nil {
		fmt.Println(n+1)
	} else {
		fmt.Println(s5, "is not an integer.")
	}
	// Output: 98

	//Use strconv.ParseInt to parse a decimal string (base 10) and check if it fits into an int64.
	s6 := "97"
	n6, err := strconv.ParseInt(s6, 10, 64)
	if err == nil {
		fmt.Printf("%d of type %T", n6, n6)
	}
	// Output: 97 of type int64


}

//Convert interface to string
func TestStringInterface(t *testing.T){
	//Use fmt.Sprintf to convert an interface value to a string.

	var x interface{} = "abc"
	str := fmt.Sprintf("%v", x)
	fmt.Println(str)

	//In fact, the same technique can be used to get a string representation of any data structure.

	var x2 interface{} = []int{1, 2, 3}
	str2 := fmt.Sprintf("%v", x2)
	fmt.Println(str2) // "[1 2 3]"
}