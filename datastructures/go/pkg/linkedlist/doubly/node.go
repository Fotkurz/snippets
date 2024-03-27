package doubly

type node[T any] struct {
	data           T
	previous, next *node[T]
}

func (n *node[T]) Data() T {
	return n.data
}

func (n *node[T]) Previous() *node[T] {
	return n.previous
}

func (n *node[T]) Next() *node[T] {
	return n.next
}
