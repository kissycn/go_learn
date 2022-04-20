package string

import "testing"

func TestByte(t *testing.T) {
	var ch1 byte = 'A'
	var ch2 byte = '\x41'
	var ch3 byte = 65

	println(ch1 == ch2 && ch1 == ch3)
}
