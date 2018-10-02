package redblack

type Tree struct {
	root *Node
	nil  *Node
}

type color int

const (
	colorRed = iota
	colorBlack
)

type Node struct {
	color color
	key   int
	left  *Node
	right *Node
	p     *Node
}

// rotate change pointer's reference to recovery red-black tree's property
func (t *Tree) LeftRotate(node *Node) {
	right := node.right
	node.right = right.left
	if right.left != t.nil {
		right.left.p = node
	}
	right.p = node.p // 替换父节点的子节点
	if right.p == nil {
		t.root = right
	} else if node == node.p.left {
		right.p.left = right
	} else {
		right.p.right = right
	}
	right.left = node
	node.p = right

}

// RightRotate node x
func (t *Tree) RightRotate(x *Node) {
	y := x.left
	x.left = y.right
	if y.right != t.nil {
		y.right.p = x
	}

	y.p = x.p
	if x.p == nil {
		t.root = y
	} else if x == x.p.right {
		x.p.right = y
	} else {
		x.p.left = y
	}

	y.right = x
	x.p = y
}

func (t *Tree) Insert(x int) {
	node := &Node{
		key:   x,
		color: colorRed,
	}
	t.InsertNode(node)
}

func (t *Tree) InsertNode(node *Node) {
	x := t.root
	y := t.nil

	for x != t.nil {
		y = x
		if node.key > x.key {
			node = node.right
		} else {
			node = node.left
		}
	}

	x.p = y
	if y == t.nil {
		t.root = x
	} else if node.key > x.key {
		x.right = node
	} else {
		x.left = node
	}

	node.left, node.right = t.nil, t.nil // 插入节点后性质被破坏
	t.fixUp(node)

}

func (t *Tree) fixUp(node *Node) {
	for node.p.color == colorRed { // 当插入结点的父节点不是红色结点时，只会破坏性质1（插入结点就是根节点）
		if node.p == node.p.p.right {
			y := node.p.p.left // 叔结点
			if y.color == colorRed {
				node.p.color = colorBlack
				y.color = colorBlack
				node.p.p.color = colorRed
				node = node.p.p
				continue
			}

			if node == node.p.right { // 插入结点为右结点，性质4被破坏，左旋转
				node = node.p
				t.LeftRotate(node)
			}

			node.p.color = colorBlack
			node.p.p.color = colorRed
			t.RightRotate(node.p.p)
		} else {
			y := node.p.p.right
			if y.color == colorRed {
				node.p.color = colorBlack
				y.color = colorBlack
				node.p.p.color = colorRed
				node = node.p
				continue
			}

			if node == node.p.right {
				node = node.p
				t.LeftRotate(node)
			}

			node.p.color = colorBlack
			node.p.p.color = colorRed
			t.RightRotate(node.p.p)

		}
	}
	t.root.color = colorBlack
}
