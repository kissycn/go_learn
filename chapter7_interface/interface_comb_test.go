package chapter7_interface

import (
	"fmt"
	"testing"
)

type Writer1 interface {
	Write(data []byte)
}

type Closer1 interface {
	Close()
}

type WriteCloser interface {
	Writer1
	Closer1
}

type Device struct {
}

func (d *Device) Write(data []byte) {
	fmt.Println(string((data)))
}

func (d *Device) Close() {
	fmt.Println("closer")
}

func TestDevice(t *testing.T) {
	var wc WriteCloser = new(Device)
	wc.Write([]byte("hello world"))
	wc.Close()

	var closer Closer1 = new(Device)
	closer.Close()
}
