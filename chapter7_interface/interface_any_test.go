package chapter7_interface

import (
	"fmt"
	"testing"
)

func TestAny(t *testing.T) {
	var any interface{}
	any = 1
	fmt.Println(any)

	any = "one"
	fmt.Println(any)

	println()
}
