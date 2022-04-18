package variable

import (
	"fmt"
	"testing"
)

// 局部变量
func TestLocalVariable(t *testing.T) {
	var a int = 1
	var b int = 2
	c := a + b
	fmt.Printf("a = %d, b = %d, c = %d", a, b, c)
}

var c int

// 全局变量
func TestGlobalVariable(t *testing.T) {
	var a int = 1
	var b int = 2
	c = a + b
	fmt.Printf("a = %d,b = %d, c = %d", a, b, c)
}

func TestFuncVariable(t *testing.T) {
	var a int = 1
	var b int = 2
	println(getData1(a, b))
}

//
func getData1(a, b int) int {
	return a + b
}
