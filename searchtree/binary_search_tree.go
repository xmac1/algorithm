package searchtree

import "fmt"

type BinarySearchTree struct {
	root *Node
}

type Node struct {
	key            int
	left, right, p *Node
}

func (t *BinarySearchTree) Root() *Node {
	return t.root
}

func (t *BinarySearchTree) InOrderWalk(node *Node) {
	if node == nil {
		return
	}

	if node.left != nil {
		t.InOrderWalk(node.left)
	}

	fmt.Println(node.key)

	if node.right != nil {
		t.InOrderWalk(node.right)
	}
}

func (t *BinarySearchTree) Minimum(root *Node) (node *Node) {
	node = root
	for node != nil {
		node = node.left
	}
	return
}

func (t *BinarySearchTree) Maximum(root *Node) (node *Node) {
	node = root
	for node != nil {
		node = node.right
	}

	return
}

func (t *BinarySearchTree) Insert(key int) {
	node := &Node{
		key: key,
	}
	t.InsertNode(node)
}

func (t *BinarySearchTree) InsertNode(node *Node) {
	root := t.root
	var parent *Node
	for root != nil {
		parent = root

		if node.key > root.key {
			root = root.right
		} else {
			root = root.left
		}
	}

	if parent == nil {
		t.root = node
		return
	}

	node.p = parent
	if node.key > parent.key {
		parent.right = node
	} else {
		parent.left = node
	}
}

func (t *BinarySearchTree) Search(key int) (node *Node) {
	root := t.root
	for root != nil {
		if root.key == key {
			node = root
			return
		}

		if root.key < key {
			root = root.right
		} else {
			root = root.left
		}
	}

	return
}

func (t *BinarySearchTree) Delete(key int) {
	node := t.Search(key)
	if node == nil {
		return
	}

	t.DeleteNode(node)
}

func (t *BinarySearchTree) DeleteNode(node *Node) {
	if node.left == nil {
		t.transplant(node, node.right)
		return
	}

	if node.right == nil {
		t.transplant(node, node.left)
		return
	}

	min := t.Minimum(node.right)
	if min != node.right {
		t.transplant(min, min.right)
		min.right = node.right
		node.right.p = min
	}
	t.transplant(node, min)
	min.left = node.left
	node.left.p = min
}

// replace node replace node's position in tree, node's parent become replace's parent
func (t *BinarySearchTree) transplant(node, replace *Node) {
	if node.p == nil {
		t.root = replace
		return
	}

	if node.p.left == node {
		node.p.left = replace
	} else {
		node.p.right = replace
	}

	if replace != nil {
		replace.p = node.p
	}
}
