package set

type HashSet[T comparable] interface {
	Add(v T) bool
	Remove(v T) bool
	Contains(v T) bool
	Slice() []T
}

type hashSet[T comparable] struct {
	m map[T]bool
}

func NewHashSet[T comparable]() HashSet[T] {
	return hashSet[T]{
		m: make(map[T]bool),
	}
}

// Add implements HashSet.
func (h hashSet[T]) Add(v T) bool {
	_, exists := h.m[v]
	h.m[v] = true

	return exists
}

// Contains implements HashSet.
func (h hashSet[T]) Contains(v T) bool {
	_, exists := h.m[v]

	return exists
}

// Remove implements HashSet.
func (h hashSet[T]) Remove(v T) bool {
	_, exists := h.m[v]
	if exists {
		delete(h.m, v)
	}

	return exists
}

func (h hashSet[T]) Slice() []T {
	var values []T
	for k := range h.m {
		values = append(values, k)
	}

	return values
}
