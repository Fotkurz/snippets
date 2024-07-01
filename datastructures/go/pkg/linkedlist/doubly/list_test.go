package doubly_test

import (
	"testing"

	"github.com/Fotkurz/datastructures/go/stack/pkg/linkedlist/doubly"
)

func TestAddNext(t *testing.T) {

	t.Run("add 1 node", func(t *testing.T) {
		list := doubly.NewDoublyLinkedList(1)

		expectedPrevious := list.Head()

		got := list.AddNext(list.Head(), 2)

		if expectedPrevious != got.Previous() {
			t.Errorf("Want != Got, want=%p, got=%p", expectedPrevious, got.Previous())
		}
		if expectedPrevious.Data() != got.Previous().Data() {
			t.Errorf("Want != Got, want=%d, got=%d", expectedPrevious.Data(), got.Previous().Data())
		}
		if got.Next() != nil {
			t.Errorf("Got next is not nil, want=nil, got=%v", got.Next())
		}
	})

	t.Run("add 2 nodes", func(t *testing.T) {
		list := doubly.NewDoublyLinkedList(1)

		second := list.AddNext(list.Head(), 2)
		third := list.AddNext(second, 3)

		if third != second.Next() {
			t.Errorf("Want != Got, want=%p, got=%p", third, second.Next())
		}
		if second != third.Previous() {
			t.Errorf("Want != Got, want=%v, got=%v", second, third.Previous())
		}
		if nil != third.Next() {
			t.Errorf("Third next is not nil, want=nil, got=%v", third.Next())
		}

	})
}

func TestAddPrevious(t *testing.T) {

	t.Run("add 1 node", func(t *testing.T) {
		list := doubly.NewDoublyLinkedList(5)

		expectedNext := list.Head()

		got := list.AddPrevious(list.Head(), 4)

		if expectedNext != got.Next() {
			t.Errorf("Want != Got, want=%p, got=%p", expectedNext, got.Next())
		}
		if expectedNext.Data() != got.Next().Data() {
			t.Errorf("Want != Got, want=%d, got=%d", expectedNext.Data(), got.Next().Data())
		}
		if got.Previous() != nil {
			t.Errorf("Got previous is not nil, want=nil, got=%v", got.Next())
		}
		if list.Head() != got {
			t.Errorf("Expected head to be = got, want=%v, got=%v", got, list.Head())
		}
	})

	t.Run("add 2 nodes and replace head", func(t *testing.T) {
		list := doubly.NewDoublyLinkedList(3)

		expectedNext := list.Head()

		second := list.AddPrevious(list.Head(), 2)
		first := list.AddPrevious(second, 1)

		if first != second.Previous() {
			t.Errorf("Want != Got, want=%p, got=%p", first, second.Previous())
		}
		if second != expectedNext.Previous() {
			t.Errorf("Want != Got, want=%v, got=%v", second, expectedNext.Previous())
		}
		if nil != first.Previous() {
			t.Errorf("Third next is not nil, want=nil, got=%v", first.Previous())
		}
		if list.Head() != first {
			t.Errorf("Expected head to be = first, want=%v, got=%v", first, list.Head())
		}
	})
}

func TestPush(t *testing.T) {

	list := doubly.NewDoublyLinkedList(1)
	list.Push(2)
	list.Push(3)

	head := list.Head()

	if head.Next().Data() != 2 {
		t.Errorf("want != got, want=%d, got=%d", 2, head.Next().Data())
	}
	if head.Next().Next().Data() != 3 {
		t.Errorf("want != got, want=%d, got=%d", 3, head.Next().Next().Data())
	}
}

func TestPop(t *testing.T) {
	list := doubly.NewDoublyLinkedList(1)
	list.Push(2)
	list.Push(3)

	list.Pop()
	head := list.Head()
	if head.Next().Next() != nil {
		t.Errorf("want != got, want=%s, got=%p", "nil", head.Next().Next())
	}

	list.Pop()
	if head.Next() != nil {
		t.Errorf("want != got, want=%s, got=%p", "nil", head.Next())
	}

}
