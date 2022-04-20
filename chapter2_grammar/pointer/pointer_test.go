package pointer

import (
	"fmt"
	"testing"
)

func TestPointer(t *testing.T) {
	var cat int = 1
	fmt.Printf("%p \n", &cat)

	var str = "Hello World"
	sp := &str

	fmt.Printf("Str type: %T \n", sp)
	fmt.Printf("Str Pointer Address: %p \n", sp)
	fmt.Printf("%s", *sp)
}
