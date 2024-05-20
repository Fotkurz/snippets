package doubly

type DoublyLinkedList[T any] struct {
	head *node[T]
}

func NewDoublyLinkedList[T any](data T) *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{
		head: &node[T]{
			data:     data,
			next:     nil,
			previous: nil,
		},
	}
}

func (l *DoublyLinkedList[T]) Head() *node[T] {
	return l.head
}

func (l *DoublyLinkedList[T]) AddNext(new T) *node[T] {
	n := node[T]{
		previous: nil,
		data:     new,
		next:     nil,
	}

	if l.head.next == nil {
		n.previous = l.head
		l.head.next = &n
	} else {
		node := l.head.next

		hasNext := true
		for hasNext {
			if node.next == nil {
				n.previous = node
				node.next = &n
				hasNext = false
			} else {
				node = node.next
			}
		}
	}

	return &n
}
