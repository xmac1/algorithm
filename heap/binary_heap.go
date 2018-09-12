package heap

type BinaryHeap struct {
	array []*Node
	idx   int
}

type Node struct {
	data int
	idx  int
}

func (n *Node) gte(other *Node) bool {
	if n.data >= other.data {
		return true
	}
	return false
}

func (n *Node) gt(other *Node) bool {
	if n.data > other.data {
		return true
	}
	return false
}

func (n *Node) lte(other *Node) bool {
	if n.data <= other.data {
		return true
	}
	return false
}

func (n *Node) lt(other *Node) bool {
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

func (bh *BinaryHeap) Parent(node *Node) *Node {
	return bh.array[node.Parent()]
}

func (bh *BinaryHeap) LeftChild(node *Node) *Node {
	if node.LeftChild() >=  len(bh.array) {
		return nil
	}
	return bh.array[node.LeftChild()]
}

func (bh *BinaryHeap) RightChild(node *Node) *Node {
	if node.RightChild() >= len(bh.array) {
		return nil
	}
	return bh.array[node.RightChild()]
}

func (hp *BinaryHeap) MaxHeapfy(i int) {
	node := hp.array[i]
	left := hp.LeftChild(node)
	right := hp.RightChild(node)

	largest := node
	if left != nil && left.gt(node) {
		largest = left
	}

	if right != nil && right.gt(largest) {
		largest = right
	}

	if largest.idx == node.idx {
		return
	}

	tmpData := node.data
	node.data = largest.data
	largest.data = tmpData
	hp.MaxHeapfy(largest.idx)
}

func (hp *BinaryHeap) MaxHeapfyLoop(i int) {
	node := hp.array[i]
	for {
		left := hp.LeftChild(node)
		right := hp.RightChild(node)
		largest := node
		if left.gt(node) {
			largest = left
		}

		if right.gt(largest) {
			largest = left
		}

		if largest.idx == node.idx {
			return
		}

		hp.array[node.idx].data = largest.data
		hp.array[largest.idx].data = node.data
		node.data = largest.data
		node.idx = largest.idx
	}
}

func (hp *BinaryHeap) InsertMaxHeap() {

}

func HeapSort(arr []int) (orderd []int) {
	hp := BuildFromArray(arr)

	orderd = make([]int, len(arr))
	for i := len(hp.array) - 1; i >= 1 ; i -- {
		orderd[i] = hp.array[0].data
		hp.array[0].data = hp.array[i].data
		hp.array = hp.array[:len(hp.array)-1]
		hp.MaxHeapfy(0)
	}

	orderd[0] = hp.array[0].data

	return
}

func BuildFromArray(arr []int) (heap *BinaryHeap) {
	bh := &BinaryHeap{
		array:make([]*Node, len(arr)),
	}
	for i := len(arr) - 1; i >=0; i-- {
		bh.array[i] = &Node{}
		bh.array[i].data = arr[i]
		bh.array[i].idx = i
	}

	for i:= len(bh.array)-1; i >= 0; i-- {
		bh.MaxHeapfy(i)
	}



	return bh
}
