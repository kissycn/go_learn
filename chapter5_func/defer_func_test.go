package chapter5_func

import (
	"fmt"
	"testing"
)

func TestDeferFunc(t *testing.T) {
	fmt.Println("Begin")
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	fmt.Println("End")
}

func foo() {
	fmt.Println("Begin")
	defer func() {
		fmt.Println("func1")
	}()
	defer func() {
		fmt.Println("func2")
	}()
	defer func() {
		fmt.Println("func3")
	}()
	fmt.Println("End")
}

func d1() int {
	var a = 1
	defer func() {
		fmt.Println()
		a = 2
	}()

	return a
}

func TestModify(f *testing.T) {
	fmt.Println(d1())
}
