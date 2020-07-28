package sequential

import "testing"

func TestRandomSelect(t *testing.T) {
	input := []int{2, 3, 4, 5, 1, 6, 8, 7}
	expect := 3
	n := RandomSelect(input, expect)
	if n != expect {
		t.Errorf("expect: %v, got: %v", expect, n)
	}
}
