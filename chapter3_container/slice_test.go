package chapter3_container

import (
	"fmt"
	"testing"
)

func TestSliceInit(t *testing.T) {
	// declare str slice
	var strList []string
	// declare int slice
	var numList []int
	// declare empty slice and init
	var numListEmpty = []int{}

	fmt.Println(strList, numList, numListEmpty)

	fmt.Println(strList == nil)
	fmt.Println(numList == nil)
	fmt.Println(numListEmpty == nil)
}

func TestMakeSlice(t *testing.T) {
	a := make([]int, 2, 4)
	b := make([]int, 5, 10)

	fmt.Println(a, b)
}

func TestAppendDataToSlice(t *testing.T) {
	arr := []int{1, 2, 3}
	arr = append(arr, 4, 5, 6)
	fmt.Println(arr)
}

func TestSliceCopy(t *testing.T) {
	slice1 := []int{1, 2, 3, 4, 5, 6}
	slice2 := []int{3, 4, 5}

	copy(slice1, slice2)
	fmt.Println(slice1)
}
