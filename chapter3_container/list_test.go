package chapter3_container

import (
	list2 "container/list"
	"fmt"
	"testing"
)

func TestListBuild(t *testing.T) {
	list := list2.New()

	list.PushBack("one")
	list.PushFront("zero")
	list.PushBack("two")
	list.PushBack("three")

	for i := list.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}
