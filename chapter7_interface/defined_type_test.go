package chapter7_interface

import (
	"fmt"
	"testing"
)

type T1 struct{}

type Interface1 interface {
	M11()
	M22()
}

func (t T1) M11() {}

func (t T1) M22() {}

type MyT T1

type MyI Interface1

func TestDefinedType(test *testing.T) {
	var t MyT
	var i MyI
	// 方法为空
	fmt.Println(t)
	// 与接口方法类型一致
	fmt.Println(i.M11, i.M22)
}
