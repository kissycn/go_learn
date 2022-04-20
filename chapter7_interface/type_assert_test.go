package chapter7_interface

import (
	"fmt"
	"testing"
)

func TestTypeAssert(t *testing.T) {
	var a interface{}
	a = 10

	val, ok := a.(int)
	if ok {
		fmt.Println(val)
	}

	var s interface{}
	s = "a"
	valS, okS := s.(int)
	if okS {
		fmt.Println(valS)
	} else {
		fmt.Println("s is now int")
	}
}

func TestGetType(t *testing.T) {
	var a interface{}
	a = 1
	getDataType(a)
}

func getDataType(t interface{}) {
	switch t.(type) {
	case int:
		fmt.Println("int value")
	case string:
		fmt.Println("str value")
	default:
		fmt.Println("defalut")
	}
}
