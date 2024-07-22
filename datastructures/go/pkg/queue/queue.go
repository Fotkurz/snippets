package queue

type Queue[T any] []T

// IsEmpty returns true if the queue is empty
func (q *Queue[T]) IsEmpty() bool {
	return len(*q) == 0
}

// Size returns the queue number of elements
func (q *Queue[T]) Size() int {
	return len(*q)
}

// Enqueue adds a new value to the end of the queue
func (q *Queue[T]) Enqueue(value T) {
	*q = append(*q, value)
}

// Dequeue removes the last element of the queue and returns it
func (q *Queue[T]) Dequeue() T {
	length := len(*q)
	if q.Size() == 0 {
		return *new(T)
	}

	last := (*q)[length-1]
	*q = (*q)[0 : length-1]
	return last
}

// Peek shows the first element of the queue
func (q *Queue[T]) Peek() *T {
	if q.Size() == 0 {
		return nil
	}

	return &(*q)[0]
}

// Back returns the last element of the queue
func (q *Queue[T]) Back() *T {
	if q.Size() == 0 {
		return nil
	}

	return &(*q)[q.Size()-1]
}
