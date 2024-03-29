package stack

type Stack[T any] []T

// Pop
//
//	Removes the top element of the stack and returns it
func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		return *new(T), false
	}

	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]

	return element, true
}

// IsEmpty
//
//	Checks if the stack is empty
func (s *Stack[T]) IsEmpty() bool {
	return len(*s) == 0
}

// Add
//
//	Adds a new element to the last position
func (s *Stack[T]) Add(value T) {
	*s = append(*s, value)
}

// Peek
//
//	Returns a copy of the top element of the stack
//	or a empty value if stack empty.
func (s *Stack[T]) Peek() T {
	copy := *s
	if len(copy)-1 < 0 {
		return *new(T)
	}

	return copy[len(copy)-1]
}

// Size
//
//	Returns the number of elements of the stack
func (s *Stack[T]) Size() int {
	return len(*s)
}
