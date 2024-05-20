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

// Head
//
//	Returns the head (first) node of the list
func (l *DoublyLinkedList[T]) Head() *node[T] {
	return l.head
}

// AddNext
//
//	Insert a new node between the actual and its next node. Returns the new inserted node.
func (l *DoublyLinkedList[T]) AddNext(actual *node[T], new T) *node[T] {
	next := actual.next
	n := &node[T]{
		data:     new,
		previous: actual,
		next:     next,
	}

	actual.next = n

	return n
}

// AddPrevious
//
//	Insert a new node between the actual and its previous node. Returns the new inserted node.
func (l *DoublyLinkedList[T]) AddPrevious(actual *node[T], new T) *node[T] {
	previous := actual.previous

	n := &node[T]{
		data:     new,
		next:     actual,
		previous: previous,
	}

	actual.previous = n

	// replaces list head for the new node
	if previous == nil {
		l.head = n
	}

	return n
}
