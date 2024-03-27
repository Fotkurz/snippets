package singly

import (
	"log"
)

type LinkedList[T any] struct {
	head *Node[T]
}

func NewLinkedList[T any](headData T) *LinkedList[T] {
	return &LinkedList[T]{
		head: &Node[T]{
			data: headData,
			next: nil,
		},
	}
}

// Add
//
//	Adds a new value to the end of the list
func (l *LinkedList[T]) Add(new T) *Node[T] {
	n := Node[T]{
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
func (l *LinkedList[T]) Head() *Node[T] {
	return l.head
}

// Length
//
//	Returns the count of items in a linked-list
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
//	Removes an element from the linked-list by index
func (l *LinkedList[T]) Delete(i int) *Node[T] {
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
