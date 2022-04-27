package chapter7_interface

import (
	"fmt"
	"testing"
)

type Interface interface {
	M1()
	M2()
}

type T struct {
	// 在结构体中嵌入接口类型
	Interface
}

func (t T) M1() {
	fmt.Println("M1")
}

func (t T) M2() {
	fmt.Println("M2")
}

func TestInterNested(test *testing.T) {
	var t T
	t.M1()
	t.M2()

	var i Interface = new(T)
	i.M1()
	i.M2()
}
