package avl

func max (x int, y int) int {
	if (x > y) {
		return x;
	} else {
		return y;
	}
}

type Node struct {
	value int
	h int 			// max subtree depth
	left *Node		// left child
	right *Node		// right child
}

type Tree struct {
	Root *Node
}

func height(node *Node) int {
	if node == nil {
		return 0
	}
	return node.h;
}

func setHeight(node *Node) {
	node.h = max(height(node.left), height(node.right)) + 1;
}

func delta(node *Node) int {
	if node == nil {
		return 0
	}
	return  height(node.left) - height(node.right)
}

func left_rot(node *Node) *Node {
	right_child := node.right
	node.right = right_child.left
	right_child.left = node
	setHeight(node)
	setHeight(right_child)
	return right_child
}

func right_rot(node *Node) *Node {
	left_child := node.left
	node.left = left_child.right
	left_child.right = node
	setHeight(node)
	setHeight(left_child)
	return left_child
}

func balance(node *Node) *Node{
	setHeight(node)
	if (delta(node) == 2) {
		if (delta(node.left) < 0) {
			node.left = left_rot(node.left)
		}
		return right_rot(node)
	}
	if (delta(node) == -2) {
		if (delta(node.right) > 0) {
			node.right = right_rot(node.right)
		}
		return left_rot(node)
	}
	return node
}

func insert(node *Node, val int) *Node {
	if node == nil {
		return &Node{value: val, left: nil, right: nil, h: 1}
	}
	if (val < node.value) {
		node.left = insert(node.left, val)
	} else {
		node.right = insert(node.right, val)
	}
	return balance(node)
}

func recursive_search(node *Node, val int) bool {
	if node == nil {
		return false
	}
	if node.value == val {
		return true
	}
	if node.value > val {
		return recursive_search(node.left, val)
	} else {
		return recursive_search(node.right, val)
	}
}



// Tree interface
func (tree *Tree) Search(value int) bool{
	return recursive_search(tree.Root, value)
}

func (tree *Tree) Insert(value int) {
	tree.Root = insert(tree.Root, value)
}
