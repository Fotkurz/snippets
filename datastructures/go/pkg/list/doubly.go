package list

type Doubly[T any] struct {
	head *doublyNode[T]
}

func NewDoubly[T any](data T) *Doubly[T] {
	l := &Doubly[T]{}

	l.head = &doublyNode[T]{
		list:     l,
		data:     data,
		next:     nil,
		previous: nil,
	}

	return l
}

// Head returns the head (first) node of the list
func (l *Doubly[T]) Head() *doublyNode[T] {
	return l.head
}

// AddNext inserts a new node between the actual and its next node.
// Returns the new inserted node.
func (l *Doubly[T]) AddNext(actual *doublyNode[T], new T) *doublyNode[T] {
	if actual.List() != l {
		return nil
	}
	next := actual.next

	n := &doublyNode[T]{
		data:     new,
		previous: actual,
		next:     next,
		list:     l,
	}

	actual.next = n

	return n
}

// AddPrevious inserts a new node between the actual and its previous node.
// Returns the new inserted node.
func (l *Doubly[T]) AddPrevious(actual *doublyNode[T], new T) *doublyNode[T] {
	if actual.List() != l {
		return nil
	}

	previous := actual.previous
	n := &doublyNode[T]{
		data:     new,
		next:     actual,
		previous: previous,
		list:     l,
	}

	actual.previous = n

	// replaces list head for the new node
	if previous == nil {
		l.head = n
	}

	return n
}

// Push adds a new node to the end of the list and returns it.
func (l *Doubly[T]) Push(data T) *doublyNode[T] {
	n := &doublyNode[T]{
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
func (l *Doubly[T]) Pop() *doublyNode[T] {
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
