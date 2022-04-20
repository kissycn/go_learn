package chapter5_func

import (
	"fmt"
	"testing"
)

func TestDeferFunc(t *testing.T) {
	fmt.Println("Begin")
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	fmt.Println("End")
}
