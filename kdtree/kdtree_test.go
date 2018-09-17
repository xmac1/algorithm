package kdtree

import (
	"fmt"
	"testing"
)

func TestMedian(t *testing.T) {
	listt := []*KDNode{
		&KDNode{X: 2, Y: 3},
		&KDNode{X: 9, Y: 6},
		&KDNode{X: 8, Y: 1},
		&KDNode{X: 4, Y: 7},
		&KDNode{X: 5, Y: 4},
		&KDNode{X: 7, Y: 2},
	}

	root := BuildKDTree(listt, 1)

	fmt.Println(root.Left)

	a := 0.0012
	b := 0.0013

	fmt.Println(a > b)
}
