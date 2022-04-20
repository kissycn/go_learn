package chapter3_container

import (
	"fmt"
	"testing"
)

func TestMapInit(t *testing.T) {
	var m map[string]int
	m = map[string]int{"one": 1, "two": 2}
	fmt.Println(m["one"])

	m2 := make(map[string]int)
	m2["one"] = 1
	fmt.Println(m2["one"])

	m3 := map[string]int{"one": 1}
	fmt.Println(m3["one"])
}

func TestMapIter(t *testing.T) {
	m1 := map[string]int{"one": 1, "two": 2}

	for k, v := range m1 {
		fmt.Println(k, v)
	}
}

func TestMapDel(t *testing.T) {
	var m1 = make(map[string]int)

	m1["one"] = 1
	m1["two"] = 2
	m1["three"] = 3

	delete(m1, "two")

	for k, v := range m1 {
		fmt.Println(k, v)
	}
}
