package convert_test

import (
	"fmt"
	"regexp"
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
}
