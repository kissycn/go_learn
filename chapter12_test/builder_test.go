package chapter12

import "testing"

func testBuilder(t *testing.T) {
	t.Log("testBuilder")
}

func testBuilderString(t *testing.T) {
	t.Log("testBuilderString")
}

func testBuilderReset(t *testing.T) {
	t.Log("testBuilderReset")
}

func testBuilderGrow(t *testing.T) {
	t.Log("testBuilderGrow")
}

func TestBuilder(t *testing.T) {
	t.Run("TestBuilder", testBuilder)
	t.Run("TestBuilderString", testBuilderString)
	t.Run("TestBuilderReset", testBuilderReset)
	t.Run("TestBuilderGrow", testBuilderGrow)
}
