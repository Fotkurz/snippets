package doubly

type DoublyLinkedList[T any] struct {
	head *node[T]
}

func NewDoublyLinkedList[T any](headData T) *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{
		head: &node[T]{
			data: headData,
			next: nil,
		},
	}
}

func (l *DoublyLinkedList[T]) Head() *node[T] {
	return l.head
}
