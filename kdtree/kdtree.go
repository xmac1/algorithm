package kdtree

import "math"

type Axis int

const (
	AxisX Axis = iota
	AxisY
)

// two dimension tree implement
type KDNode struct {
	X, Y                float32
	Dist                float32
	Left, Right, Parent *KDNode
	Country             string
	Axis                Axis
}

func (node *KDNode) distance(x, y float32) float32 {
	return float32(math.Sqrt(math.Pow(float64(node.X)-float64(x), 2) + math.Pow(float64(node.Y)-float64(y), 2)))
}

func (node *KDNode) distanceNode(n *KDNode) float32 {
	return node.distance(n.X, n.Y)
}

func (node *KDNode) SearchNearest(x, y float32) *KDNode {
	almostNode := node.SearchLeaf(x, y)

	for {
		if almostNode == nil {
			break
		}

		almostNode.Dist = almostNode.distance(x, y)
		dist := float64(almostNode.Dist)

		if almostNode.Parent != nil {
			var parent bool
			if almostNode.Parent.Axis == AxisX {
				if dist > math.Abs(float64(x-almostNode.Parent.X)) {
					parent = true
				}
			} else {
				if dist > math.Abs(float64(y-almostNode.Parent.Y)) {
					parent = true
				}
			}

			var brother *KDNode
			if parent {
				if almostNode == almostNode.Parent.Left {
					brother = almostNode.Parent.Right
				} else {
					brother = almostNode.Parent.Left
				}

			} else {
				break
			}

			if brother != nil {
				brother.Dist = brother.distance(x, y)
				if brother.Dist < almostNode.Dist {
					almostNode = brother
				} else {
					almostNode = almostNode.Parent
				}
			}

		}
	}
	return node
}

func (node *KDNode) SearchDepthFirst(x, y float32) (nearestNode *KDNode) {
	nearestNode = node

	dist := nearestNode.distance(x, y)
	left := nearestNode.Left
	if left != nil {
		lDist := left.distance(x, y)
		if lDist < dist {
			nearestNode = left
			nearestNode = nearestNode.SearchDepthFirst(x, y)
		}
	}

	dist = nearestNode.distance(x, y)

	right := node.Right
	if right != nil {
		rDist := right.distance(x, y)
		if rDist < dist {
			nearestNode = right
			nearestNode = nearestNode.SearchDepthFirst(x, y)
		}
	}

	return

}

func (node *KDNode) SearchLeaf(x, y float32) *KDNode {
	leaf := node
	var next *KDNode

	axis := AxisX

	for {
		if leaf.Left == nil && leaf.Right == nil {
			break
		}

		var (
			in, leafValue float32
		)

		switch axis {
		case AxisX:
			in = x
			leafValue = leaf.X
		case AxisY:
			in = y
			leafValue = leaf.Y
		}

		if in < leafValue {
			next = leaf.Left
		} else if in > leafValue {
			next = leaf.Right
		} else {
			if leaf.Left == nil {
				next = leaf.Right
			} else if leaf.Right == nil {
				next = leaf.Left
			} else if leaf.Left.distance(x, y) > leaf.Left.distance(x, y) {
				next = leaf.Left
			} else {
				next = leaf.Right
			}
		}

		if next == nil {
			break
		}

		leaf = next
		axis = (axis + 10%2) + 1
	}

	return leaf

}

func BuildKDTree(nodeList []*KDNode, dimension Axis) *KDNode {
	if len(nodeList) <= 0 {
		return nil
	}

	if len(nodeList) == 1 {
		return nodeList[0]
	}

	m := quickMedian(nodeList, dimension)

	node := nodeList[m]
	node.Axis = dimension

	leftNodeList := nodeList[:m]
	rightNodeList := nodeList[m+1:]

	dimension = (dimension + 1) % 2

	node.Left = BuildKDTree(leftNodeList, dimension)
	node.Right = BuildKDTree(rightNodeList, dimension)

	if node.Left != nil {
		node.Left.Parent = node
	}
	if node.Right != nil {
		node.Right.Parent = node
	}

	return node

}

func quickMedian(list []*KDNode, x Axis) int {
	if len(list) <= 1 {
		return 0
	}

	left := 0
	right := len(list)

	half := len(list) >> 1

	var m int
	for {
		m = left + partition(list[left:right], x)
		if m == half {
			break
		}

		if m > half {
			right = m
		} else {
			left = m + 1
		}
	}

	return m

}

func partition(list []*KDNode, x Axis) int {
	if len(list) <= 1 {
		return 0
	}

	node := list[len(list)-1]

	var i, j = 0, -1
	for ; i < len(list)-1; i++ {
		n := list[i]
		if x == AxisX {
			if n.X < node.X {
				j++
				if i != j {
					tmp := list[j]
					list[j] = list[i]
					list[i] = tmp
				}
			}
		} else {
			if n.Y < node.Y {
				j++
				if i != j {
					tmp := list[j]
					list[j] = list[i]
					list[i] = tmp
				}
			}
		}
	}

	j++
	tmp := list[j]
	list[j] = node
	list[len(list)-1] = tmp
	return j
}
