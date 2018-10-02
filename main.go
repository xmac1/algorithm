package main

import "github.com/xmac1/example/searchtree"

func main() {
	arr := []int{5, 3, 4, 1, 6, 2}

	t := &searchtree.BinarySearchTree{}

	for _, key := range arr {
		t.Insert(key)
	}

	t.Insert(7)

	t.Delete(1)
	t.Delete(6)
	t.InOrderWalk(t.Root())
}
