package tree

type Node struct {
	Parent *Node
	Left   *Node
	Right  *Node
}

func NewNode(parent *Node) *Node {
	n := &Node{}

	if parent != nil {
		n.Parent = parent
	}

	return n
}

func (n *Node) IsLeaf() bool {
	return n.Left == nil && n.Right == nil
}

func (n *Node) Root() *Node {
	node := n
	for node.Parent != nil {
		node = node.Parent
	}
	return node
}

func Height(n *Node) uint {
	if n == nil {
		return 0
	}

	// get to the root
	var root *Node
	if n.Parent != nil {
		root = n.Parent
	} else {
		root = n
	}

	for root != nil {
		root = n.Parent
	}

	// get to the leafs
	var height uint = 1
	node := root
	for {
		if node.Left == nil && node.Right != nil {
			height += 1
			node = node.Right
		}
		if node.Right == nil && node.Left != nil {
			height += 1
			node = node.Left
		}
		if node.IsLeaf() {
			break
		}
	}

	return height
}
