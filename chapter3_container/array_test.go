package chapter3_container

import (
	"fmt"
	"testing"
)

func TestArrayInit(t *testing.T) {
	var a1 [3]int
	a1[0] = 1
	a1[1] = 2
	a1[2] = 3

	println(a1[2])

	for i, v := range a1 {
		fmt.Printf("index:%d value:%d \n", i, v)
	}

	fmt.Println(a1)
}

func TestArrayInit1(t *testing.T) {
	a1 := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	println(a1[0])

	a2 := [...]int{4, 5, 6, 7, 8, 9}
	println(a2[2])

	a3 := [2]string{"hello", "world"}
	println(a3[1])
}

func TestArraySub(t *testing.T) {
	a1 := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	a2 := a1[0:2]
	println(a2[1])

	a3 := [4]string{"hello", "world", "hello", "go"}
	a4 := a3[2:]
	for _, v := range a4 {
		println(v)
	}
}
