package main

import (
	"fmt"

	"github.com/xmac1/example/kdtree"
)

func main() {
	nodeList := []*kdtree.KDNode{
		&kdtree.KDNode{X: 2, Y: 3},
		&kdtree.KDNode{X: 4, Y: 5},
		&kdtree.KDNode{X: 5, Y: 4},
		&kdtree.KDNode{X: 7, Y: 2},
		&kdtree.KDNode{X: 8, Y: 1},
		&kdtree.KDNode{X: 9, Y: 6},
	}

	root := kdtree.BuildKDTree(nodeList, kdtree.AxisX)

	node := root.SearchDepthFirst(0, 0)
	fmt.Println(node)
}
