package linkedlist

type LinkedList[T any] struct {
	head *Node[T]
}

func NewLinkedList[T any](data T) *LinkedList[T] {
	return &LinkedList[T]{
		head: &Node[T]{
			data: data,
			next: nil,
		},
	}
}

// Add
//
//	Adds a new value to the end of the list
func (l *LinkedList[T]) Add(new T) *Node[T] {
	n := &Node[T]{
		data: new,
		next: nil,
	}
	// TODO: Not properly working i guess
	if l.head.next == nil {
		l.head.next = n
	} else {
		node := l.head.next

		hasNext := true
		for hasNext {
			if node.next == nil {
				node.next = n
				hasNext = false
			} else {
				node = node.next
			}
		}

	}

	return n
}

func (l *LinkedList[T]) Head() *Node[T] {
	return l.head
}
