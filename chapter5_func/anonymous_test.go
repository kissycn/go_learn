package chapter5_func

import (
	"fmt"
	"testing"
)

func TestAnonymous(t *testing.T) {
	func(data string) {
		fmt.Println("Hello", data)
	}("world")
}

// 函数作为类型（有点类似于接口的概念）
type PrinterFunc func(s string)

func Print(s string) {
	fmt.Println(s)
}

func TestFuncType(t *testing.T) {
	// 使用函数并声明实现
	var printer PrinterFunc = Print
	printer("test")
}

func TestFuncVariable(t *testing.T) {
	var f = Print
	f("test")
}

func TestFunc(t *testing.T) {
	f := Hello()
	s := f("World")
	fmt.Println(s)
}

// 定义一个 func(s string) string 为返回值的函数
func Hello() func(s string) string {
	return func(s string) string {
		return "Hello " + s
	}
}
