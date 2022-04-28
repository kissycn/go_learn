package chapter12

import "testing"

func Add(a, b int) int {
	return b + a
}

func TestAdd(t *testing.T) {
	var arrs = []struct {
		a, b, res int
	}{
		{1, 2, 3},
		{3, 5, 8},
		{0, 0, 0},
	}

	for _, d := range arrs {
		i := Add(d.a, d.b)
		if i != d.res {
			t.Errorf("want %d but add(%d,%d) = %d", d.res, d.a, d.b, i)
		}
	}
}
