package linkedlist

import (
	"fmt"
	"testing"
)

func TestFront(t *testing.T) {
	l := New()
	l.PushFront(3)
	l.PushFront(2)
	l.PushFront(1)

	fmt.Println(l.len)
	fmt.Println(l.Front().Value, "---", l.Back().Value)
}
