package chapter7_interface

import (
	"fmt"
	"testing"
)

type Flyer interface {
	Fly()
}

type Walker interface {
	Walk()
}

type Bird struct {
}

func (Bird) Fly() {
	fmt.Println("bird fly")
}

func (Bird) Walk() {
	fmt.Println("bird walk")
}

type Pig struct {
}

func (Pig) Walk() {
	fmt.Println("pig walk")
}

func TypeAssert(v interface{}) {
	switch v.(type) {
	case Walker:
		fmt.Printf("walker")
	case Flyer:
		fmt.Printf("flyeer")
	default:
		fmt.Printf("unknow")
	}
}

func TestTypeAssert1(t *testing.T) {
	var w Walker = new(Pig)
	TypeAssert(w)
}

func TestInterfaceConvert(t *testing.T) {
	animals := map[string]interface{}{
		"pig":  new(Pig),
		"bird": new(Bird),
	}

	for _, obj := range animals {
		f, isFlyer := obj.(Flyer)
		w, isWalker := obj.(Walker)

		if isFlyer {
			f.Fly()
		}

		if isWalker {
			w.Walk()
		}
	}

	var w Walker = new(Pig)
	obj, ok := w.(Walker)
	if ok {
		obj.Walk()
	}

	var p1 = new(Pig)
	var ww Walker
	ww = p1
	obj1, ok1 := ww.(Walker)
	if ok1 {
		obj1.Walk()
	}
}
