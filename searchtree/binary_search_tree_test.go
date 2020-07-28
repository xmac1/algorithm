package searchtree

import (
	"reflect"
	"testing"
)

func TestBinarySearchTree_InOrderWalk(t *testing.T) {
	output := make([]int, 0, 10)
	tree := New(func(node *Node) {
		output = append(output, node.key)
	})
	tree.Insert(3)
	tree.Insert(5)
	tree.Insert(2)
	tree.Insert(1)
	tree.Insert(4)
	tree.InOrderWalk(tree.Root())
	expect := []int{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(output, expect) {
		t.Errorf("in-order walk failed, want: %v, get: %v", expect, output)
	}
}
