package string

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	var s string = "Hello \n GoLang"
	println(s)
	println("len = ", len(s))
	println(string(s[0]), s[8])

	var s1 string = "Hello "
	var s2 string = "World"
	s3 := s1 + s2
	println(s3)

	const s4 = `num1,
num2,
num3`

	println(s4)
}

func TestStringIter(t *testing.T) {
	var str string = "Hello World"

	for _, s := range str {
		fmt.Printf("%c \n", s)
	}
}
