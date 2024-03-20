package linkedlist

import "fmt"

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
	n := &Node[T]{
		data: new,
		next: nil,
	}

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
	fmt.Println(&n)
	return n
}

func (l *LinkedList[T]) Head() *Node[T] {
	return l.head
}
