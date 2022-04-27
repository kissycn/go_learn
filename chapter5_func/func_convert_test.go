package chapter5_func

import (
	"fmt"
	"testing"
)

type BinaryAdder interface {
	Add(int, int) int
}

type MyAdderFunc func(int, int) int

func (f MyAdderFunc) Add(x, y int) int {
	return f(x, y)
}

func MyAdd(x, y int) int {
	return x + y
}

func TestFuncConvert(t *testing.T) {
	var b BinaryAdder = MyAdderFunc(MyAdd)
	fmt.Println(b.Add(1, 2))
}
