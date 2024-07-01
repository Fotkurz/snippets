package singly

import (
	"log"
)

type LinkedList[T any] struct {
	head *node[T]
}

func NewLinkedList[T any](data T) *LinkedList[T] {
	return &LinkedList[T]{
		head: &node[T]{
			data: data,
			next: nil,
		},
	}
}

// Add
//
//	Adds a new value to the end of the list and return its node.
func (l *LinkedList[T]) Add(new T) *node[T] {
	n := node[T]{
		data: new,
		next: nil,
	}

	if l.head.next == nil {
		l.head.next = &n
	} else {
		node := l.head.next

		hasNext := true
		for hasNext {
			if node.next == nil {
				node.next = &n
				hasNext = false
			} else {
				node = node.next
			}
		}

	}

	return &n
}

// Head
//
//	Returns the head node from the list
func (l *LinkedList[T]) Head() *node[T] {
	return l.head
}

// Length
//
//	Returns the count of items of the list
func (l *LinkedList[T]) Length() int {
	node := l.Head()
	count := int(1)

	for {
		if node.Next() != nil {
			node = node.Next()
			count++
		} else {
			break
		}
	}

	return count
}

// Delete
//
//	Removes an element from the list by index
func (l *LinkedList[T]) Delete(i int) *node[T] {
	if i > l.Length()-1 || i < 0 {
		log.Panicf("index out of bounds [%d]", i)
	}

	node := l.Head()
	if i == 0 {
		l.head = node.Next()
	} else {
		before := node
		for n := 0; n < i; n++ {
			before = node
			node = node.Next()
		}

		before.next = node.Next()
	}

	return node
}
