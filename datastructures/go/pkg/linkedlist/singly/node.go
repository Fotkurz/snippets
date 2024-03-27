package singly

type Node[T any] struct {
	data T
	next *Node[T]
}

// Data
//
//	Returns this node data.
func (n *Node[T]) Data() T {
	return n.data
}

// Next
//
//	Returns the next node.
func (n *Node[T]) Next() *Node[T] {
	return n.next
}
