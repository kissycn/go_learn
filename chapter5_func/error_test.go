package chapter5_func

import (
	"errors"
	"fmt"
	"testing"
)

var errDivisionByZero = errors.New("division by zero\n")

func div(num1, num2 int) (int, error) {
	if num2 == 0 {
		return 0, errDivisionByZero
	}

	return num1 / num2, nil
}

func TestErr(t *testing.T) {
	_, err := div(1, 0)
	if nil != err {
		fmt.Println(err)
	}
	fmt.Println(div(4, 2))
}
