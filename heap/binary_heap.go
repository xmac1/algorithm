package heap

type BinaryHeap struct {
	array []Node
	idx   int
}

type Node struct {
	data int
	idx  int
}

func (n *Node) gte(other Node) bool {
	if n.data >= other.data {
		return true
	}
	return false
}

func (n *Node) gt(other Node) bool {
	if n.data > other.data {
		return true
	}
	return false
}

func (n *Node) lte(other Node) bool {
	if n.data <= other.data {
		return true
	}
	return false
}

func (n *Node) lt(other Node) bool {
	if n.data < other.data {
		return true
	}
	return false
}

func (n *Node) Parent() int {
	return n.idx / 2
}

func (n *Node) LeftChild() int {
	return n.idx*2 + 1
}

func (n *Node) RightChild() int {
	return (n.idx + 1) * 2
}

func (bh *BinaryHeap) Parent(node Node) Node {
	return bh.array[node.Parent()]
}

func (bh *BinaryHeap) LeftChild(node Node) Node {
	return bh.array[node.LeftChild()]
}

func (bh *BinaryHeap) RightChild(node Node) Node {
	return bh.array[node.RightChild()]
}

func (hp *BinaryHeap) MaxHeapfy(i int) {
	node := hp.array[i]
	left := hp.LeftChild(node)
	right := hp.RightChild(node)

	var largest Node
	if left.gt(node) {
		largest = left
	}

	if right.gt(largest) {
		largest = right
	}

	if largest.idx == node.idx {
		return
	}

	hp.array[node.idx].data = largest.data
	hp.array[largest.idx].data = node.data
	hp.MaxHeapfy(largest.idx)
}

func (hp *BinaryHeap) InsertMaxHeap() {

}

func BuildFromArray(arr []int) (heap *BinaryHeap) {

}
