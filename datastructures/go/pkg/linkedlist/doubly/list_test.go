package doubly_test

import (
	"testing"

	"github.com/Fotkurz/datastructures/go/stack/pkg/linkedlist/doubly"
)

func TestAdd(t *testing.T) {

	t.Run("add 1 node", func(t *testing.T) {
		list := doubly.NewDoublyLinkedList[int](5)

		nextOfHead := 6
		previousOfFirstNode := 5

		got := list.AddNext(6)

		if nextOfHead != got.Data() {
			t.Errorf("Want != Got, want=%d, got=%d", nextOfHead, got.Data())
		}
		if list.Head().Previous() != nil {
			t.Errorf("Head previous node is not nil, got=%v", list.Head().Previous())
		}
		if list.Head().Data() != previousOfFirstNode {
			t.Errorf("Want != Got, want=%v, got=%v", list.Head(), got.Previous())
		}

	})
}
