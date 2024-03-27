package doubly

type Node[T any] struct {
	data           T
	previous, next *Node[T]
}
