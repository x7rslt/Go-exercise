package switch_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

//Basic switch with default
func TestBasicSwitch(t *testing.T){
	switch time.Now().Weekday() {
	case time.Saturday:
		fmt.Println("Today is Saturday.")
	case time.Sunday:
		fmt.Println("Today is Sunday.")
	default:
		fmt.Println("Today is a weekday.")
	}
}

//No condition
func TestNoCondition(t *testing.T){
	switch hour := time.Now().Hour(); { // missing expression means "true"
	case hour < 12:
		fmt.Println("Good morning!")
	case hour < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}
}


func WhiteSpace(c rune) bool {
	switch c {
	case ' ', '\t', '\n', '\f', '\r':
		return true
	}
	return false
}
//Case list
func TestCaseList(t *testing.T){
	WhiteSpace(' ')
}

//Fallthrough
//A fallthrough statement transfers control to the next case.
//It may be used only as the final statement in a clause.
func TestFallthrough(t *testing.T){
	switch 1 {
	case 1:
		fmt.Println("1")
		fallthrough
	case 2:
		fmt.Println("2")
		fallthrough
	case 3:
		fmt.Println("3")
	}
}
//Exit with break
//A break statement terminates execution of the innermost for, switch, or select statement.
//
//If you need to break out of a surrounding loop, not the switch, you can put a label on the loop and break to that label. This example shows both uses.
func TestBreakSwitch(t *testing.T){
Loop:
	for _, ch := range "a b\nc" {
		switch ch {
		case ' ': // skip space
			break
		case '\n': // break at newline
			break Loop
		default:
			fmt.Printf("%c\n", ch)
		}
	}
}

//Execution order
//First the switch expression is evaluated once.
//Then case expressions are evaluated left-to-right and top-to-bottom:
//the first one that equals the switch expression triggers execution of the statements of the associated case,
//the other cases are skipped.
// Foo prints and returns n.
func Foo(n int) int {
	fmt.Println(n)
	return n
}

func TestExecution(t *testing.T){
	switch Foo(2) {
	case Foo(1), Foo(2), Foo(3):
		fmt.Println("First case")
		fallthrough
	case Foo(4):
		fmt.Println("Second case")
	}
}

func TestSwitch(t *testing.T) {
	rand.Seed(time.Now().UnixNano()) //生成随机数
	switch num := rand.Intn(30); num {
	case 1:
		fmt.Println("this is 1")
	case 2:
		fmt.Println("this is 2")
	default:
		fmt.Println("this is no you choice.")
	}
}

func TestSwitch2(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	switch num := rand.Intn(100); {
	case num < 10:
		fmt.Println(1)
	case num == 11:
		fmt.Println(11)
	default:
		fmt.Println("no rand number as your wish.")
	}
}

func TestSwitch3(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	num := rand.Intn(100)
	switch {
	case num < 10:
		fmt.Println("num is low")
	case num > 50:
		fmt.Println("num is big")
	default:
		fmt.Println("num is good.")
	}
}
