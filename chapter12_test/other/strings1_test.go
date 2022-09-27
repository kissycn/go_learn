package other_test

import (
	"fmt"
	"testing"
)

func setUp1() func() {
	fmt.Println("setup start...")
	return func() {
		fmt.Printf("setup tearDown")
	}
}

func TestCompare1(t *testing.T) {
	t.Cleanup(setUp1())
	t.Log("TestCompare")
}

func TestCompareStrings1(t *testing.T) {
	t.Cleanup(setUp1())
	t.Log("TestCompareStrings")
}

func pkgSetUp() func() {
	fmt.Println("pkg setup start...")
	return func() {
		fmt.Printf("pkg setup tearDown")
	}
}

//func TestMain(m *testing.M)  {
//	defer pkgSetUp()()
//	m.Run()
//}
