package quicksort

import (
	"reflect"
	"testing"
)

func TestMedium(t *testing.T) {
	input := []int{4, 2, 1, 3, 5}
	Sort(input)
	expect := []int{1, 2, 3 ,4, 5}
	if !reflect.DeepEqual(input, expect) {
		t.Errorf("quick sort failed, want: %v, got: %v", expect, input)
	}
}
