package doubly

type DoublyLinkedList[T any] struct {
	head *Node[T]
}

func NewDoublyLinkedList[T any](headData T) *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{
		head: &Node[T]{
			data: headData,
			next: nil,
		},
	}
}

func (l *DoublyLinkedList[T]) Head() *Node[T] {
	return l.head
}
