package interface_test

import (
	"fmt"
	"testing"
)

type error interface {
	Error() string
}

type networkProblem struct {
	message string
	code    int
}

func (np networkProblem) Error() string {
	return fmt.Sprintf("network error! message:%s,code:%v", np.message, np.code)
}

func TestError(t *testing.T) {
	//sanonymous function
	handleError := func(err error) {
		fmt.Println(err.Error())
	}
	np := networkProblem{
		message: "we received a problem",
		code:    404,
	}
	handleError(np)
}
