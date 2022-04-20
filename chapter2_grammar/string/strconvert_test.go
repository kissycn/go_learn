package string

import (
	"fmt"
	"strconv"
	"testing"
)

func TestConvertIntStr(t *testing.T) {
	var a int = 1
	b := strconv.Itoa(a)
	fmt.Printf("type: %T \nvalue: %v \n", b, b)
}

func TestConvertStrInt(t *testing.T) {
	var s string = "100"
	i, _ := strconv.Atoi(s)
	fmt.Printf("type: %T \n value:%v \n", i, i)
}
