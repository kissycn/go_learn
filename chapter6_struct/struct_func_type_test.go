package chapter6_struct

import (
	"fmt"
	"testing"
)

type MyInt int

func (i MyInt) IsZero() bool {
	return i == 0
}

func (i MyInt) Add(num int) int {
	return int(i) + num
}

func TestMyInt(t *testing.T) {
	var i MyInt = 1
	if !i.IsZero() {
		fmt.Println(i)
	}
}

func TestMyIntAdd(t *testing.T) {
	var i MyInt = 1
	res := i.Add(2)
	fmt.Println(res)
}
