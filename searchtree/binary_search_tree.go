package searchtree

import "fmt"

type BinarySearchTree struct {
	left  *BinarySearchTree
	right *BinarySearchTree
	data  int
}

func InOrderTreeWalk(tree *BinarySearchTree) {
	if tree == nil {
		return
	}

	if tree.left != nil {
		InOrderTreeWalk(tree.left)
	}

	fmt.Println(tree.data)

	if tree.right != nil {
		InOrderTreeWalk(tree.right)
	}
}

func Search(tree *BinarySearchTree, key int) (node *BinarySearchTree) {
	if tree == nil {
		return
	}

	if tree.data == key {
		return tree
	}

	if tree.data < key {
		return Search(tree.right, key)
	}

	return Search(tree.left, key)
}

func Maximum(tree *BinarySearchTree) (node *BinarySearchTree) {
	if tree == nil {
		return nil
	}
	for tree.right != nil {
		tree = tree.right
	}

	node = tree
	return
}

func Minimum(tree *BinarySearchTree) (node *BinarySearchTree) {
	if tree.left == nil {
		return
	}

	for tree.left != nil {
		tree = tree.left
	}

	node = tree
	return
}

func Insert(tree *BinarySearchTree, data int) (node *BinarySearchTree) {
	if tree == nil {
		node = &BinarySearchTree{
			left:  nil,
			right: nil,
			data:  data,
		}
		return
	}

}
