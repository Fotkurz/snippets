package list

type doublyNode[T any] struct {
	data           T
	previous, next *doublyNode[T]
	list           *Doubly[T]
}

func (n *doublyNode[T]) Data() T {
	return n.data
}

func (n *doublyNode[T]) Previous() *doublyNode[T] {
	return n.previous
}

func (n *doublyNode[T]) Next() *doublyNode[T] {
	return n.next
}

func (n *doublyNode[T]) HasNext() bool {
	return n.next != nil
}

func (n *doublyNode[T]) List() *Doubly[T] {
	return n.list
}

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
