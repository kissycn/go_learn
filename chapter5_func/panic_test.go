package chapter5_func

import (
	"fmt"
	"testing"
)

func TestPanic(t *testing.T) {
	defer fmt.Println("before panic")
	defer fmt.Println("before panic")

	panic("panic")
}

func TestPanicRecover(T *testing.T) {
	fmt.Println("Begin")

	defer func() {
		err := recover()
		fmt.Println("======", err)
		fmt.Println("End")
	}()

	panic("panic error")
}
