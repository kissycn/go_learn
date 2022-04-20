package variable

import (
	"fmt"
	"testing"
)

type IntAlias = int
type NewInt int

func TestTypeAlias(t *testing.T) {
	var a IntAlias = 1
	var b NewInt = 1
	var c int = 1

	fmt.Printf("type alias is:%T \n  new type is:%T \n int type is:%T", a, b, c)
}
