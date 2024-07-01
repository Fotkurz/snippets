package singly

type node[T any] struct {
	data T
	next *node[T]
}

// Data
//
//	Returns this node data.
func (n *node[T]) Data() T {
	return n.data
}

// Next
//
//	Returns the next node.
func (n *node[T]) Next() *node[T] {
	return n.next
}
