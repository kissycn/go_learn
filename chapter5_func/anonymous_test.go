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
