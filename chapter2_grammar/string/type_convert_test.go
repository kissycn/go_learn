package string

import "testing"

func TestTypeConvert(t *testing.T) {
	var f1 float32 = 5.23
	f2 := int(f1)

	println(f2)
}
