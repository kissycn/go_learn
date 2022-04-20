package chapter6_struct

import (
	"fmt"
	"testing"
)

// 定义点结构
type Point struct {
	X int
	Y int
}

// 非指针接收器的加方法
func (p Point) Add(other Point) Point {
	// 成员值与参数相加后返回新的结构
	return Point{p.X + other.X, p.Y + other.Y}
}

func TestStructPoint(t *testing.T) {
	p1 := Point{1, 1}
	p2 := Point{2, 2}
	p3 := p1.Add(p2)

	fmt.Println(p3)
}

type Property struct {
	name  string
	value int
}

func (p Property) SetValue() {
	fmt.Println("sub func", p)
	p.name = "two"
	p.value = 2
	fmt.Println("sub func", p)
}

func (p *Property) SetValue1() {
	p.name = "two"
	p.value = 2
}

func TestSetValue(t *testing.T) {
	var p *Property = new(Property)
	p.name = "one"
	p.value = 1
	p.SetValue()
	fmt.Println("main func", p)

	p.name = "three"
	p.value = 3
	p.SetValue1()
	fmt.Println("reference func", p)
}
