package main

import (
	"container/list"
	"fmt"
)

func main() {
	code := list.New()
	code.PushBack("code")

	fmt.Println(code.Front().Value)
}
