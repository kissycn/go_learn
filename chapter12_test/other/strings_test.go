package other_test

import (
	"fmt"
	"testing"
)

func setUp() func() {
	fmt.Println("setup start...")
	return func() {
		fmt.Printf("setup tearDown")
	}
}

func TestCompare(t *testing.T) {
	t.Cleanup(setUp())
	t.Log("TestCompare")
}

func TestCompareStrings(t *testing.T) {
	t.Cleanup(setUp())
	t.Log("TestCompareStrings")
}

func TestCompareIdenticalStrings(t *testing.T) {
	t.Cleanup(setUp())
	t.Log("TestCompareIdenticalStrings")
}

func TestBuilder(t *testing.T) {
	t.Cleanup(setUp())
	t.Log("TestBuilder")
}

func TestBuilderString(t *testing.T) {
	t.Cleanup(setUp())
	t.Log("TestBuilderString")
}

func TestBuilderReset(t *testing.T) {
	t.Cleanup(setUp())
	t.Log("TestBuilderReset")
}

func TestBuilderGrow(t *testing.T) {
	t.Cleanup(setUp())
	t.Log("TestBuilderGrow")
}
