package variable

import "testing"

/**
变量声明
*/
func TestVariableDeclare(t *testing.T) {
	// single declare
	var a int
	a = 1
	println(a)
	// batch declare
	var c, d float32
	c = 1.1
	d = 2
	println(c, d)
	// cur short declare and init
	f := 123.2
	g := 1
	println(f, g)
}

func TestVariableInit(t *testing.T) {
	var hd int = 10
	println(hd)

	var a = 10
	var b = 20
	println(a + b)

	var f bool = false
	var f1 bool = true

	println(f == f1)

	f3 := false
	f4 := false
	println(f4 == f3)
}

func TestVariableExchange(t *testing.T) {
	var a int = 10
	b := 20

	b, a = a, b
	println(a, b)
}

// 匿名变量
func TestVariableAnonymous(t *testing.T) {
	a, _ := getData()
	_, b := getData()

	println(a, b)
}

func getData() (int, int) {
	return 10, 20
}
