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

// Head returns the head (first) node of the list
func (l *DoublyLinkedList[T]) Head() *node[T] {
	return l.head
}

// AddNext inserts a new node between the actual and its next node.
// Returns the new inserted node.
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

// AddPrevious inserts a new node between the actual and its previous node.
// Returns the new inserted node.
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

// Push adds a new node to the end of the list and returns it.
func (l *DoublyLinkedList[T]) Push(data T) *node[T] {
	n := &node[T]{
		data:     data,
		previous: nil,
		next:     nil,
	}

	node := l.head
	hasNext := true
	for hasNext {
		if node.HasNext() {
			node = node.next
			continue
		}

		hasNext = false
		node.next = n
		n.previous = node
	}

	return n
}

// Pop removes the last item of the list and returns it.
// If the list has only head node
func (l *DoublyLinkedList[T]) Pop() *node[T] {
	n := l.head
	node := l.head

	hasNext := true
	for hasNext {
		if node.HasNext() {
			node = node.next
			continue
		}

		hasNext = false

		prev := node.previous
		prev.next = nil

		n = node
	}

	return n
}
