package switch_test

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestOne(t *testing.T) {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	default:
		fmt.Println("This is a computer.")
	}
}

func TestOrder(t *testing.T) {
	//Switch evaluation order
	//Switch cases evaluate cases from top to bottom, stopping when a case succeeds.
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	fmt.Println(time.Saturday, today)
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tommorrow")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("too far away.")
	}
}

func TestNoCondition(t *testing.T) {
	//Switch without a condition is the same as switch true.
	n := time.Now()
	switch {
	case n.Hour() < 12:
		fmt.Println("Good morning.")
	case n.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
