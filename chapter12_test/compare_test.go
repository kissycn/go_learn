package chapter12

import (
	"fmt"
	"testing"
)

func testCompare(t *testing.T) {
	t.Log("testCompare")
}

func testCompareIdenticalString(t *testing.T) {
	t.Log("testCompareIdenticalString")
}

func testCompareStrings(t *testing.T) {
	t.Log("testCompareStrings")
}

func suitSetUp() func() {
	fmt.Println("suit setup start")
	return func() {
		fmt.Printf("suit teardown")
	}
}

func suitPkgSetUp() func() {
	fmt.Println("pkg suit setup start")
	return func() {
		fmt.Printf("pkg suit teardown")
	}
}

func TestMain(m *testing.M) {
	defer suitPkgSetUp()()
	m.Run()
}

func TestCompare1(t *testing.T) {
	t.Cleanup(suitSetUp())
	t.Run("Compare", testCompare)
	t.Run("CompareString", testCompareStrings)
	t.Run("CompareIdenticalString", testCompareIdenticalString)
}
