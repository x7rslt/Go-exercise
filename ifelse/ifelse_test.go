package ifelse_test

import (
	"fmt"
	"testing"
)

func TestIf(t *testing.T) {
	a := 10
	if a > 0 {
		fmt.Println("a big than 0")
	} else {
		fmt.Println("a small than 0")
	}
}

func TestElif(t *testing.T) {
	a := 10
	if a > 5 {
		fmt.Println("a big than 5")
		if a > 12 {
			fmt.Println("a big than 12")
		} else if a > 10 {
			fmt.Println("a big than 10")
		} else if a < 10 {
			fmt.Println("a big than 10")
		} else {
			fmt.Println("a is 10")
		}
	}
}


//Loop
func TestLoop(t *testing.T){
	sum := 0
	for i:= 1;i<5;i++{
		if i%2!=0{
			continue
		}else if i == 4 {
			fmt.Println("i reach break")
			return
		}
		//因为上面的return下面不会执行
		sum +=i
		fmt.Println(sum)
	}

}
