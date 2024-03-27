package singly_test

import (
	"testing"

	"github.com/Fotkurz/datastructures/go/stack/pkg/linkedlist/singly"
)

func TestAdd(t *testing.T) {
	headData := 12
	l := singly.NewLinkedList(headData)

	l.Add(16)
	third := l.Add(17)
	fourth := l.Add(18)

	headNode := l.Head()
	thirdNode := headNode.Next().Next()
	fourthNode := thirdNode.Next()

	if headNode.Data() != headData {
		t.Errorf("Expected=%d, got=%d", headData, headNode.Data())
	}

	if third.Data() != thirdNode.Data() {
		t.Errorf("Expected=%d, got=%d", thirdNode.Data(), third.Data())
	}

	if third != thirdNode {
		t.Errorf("Expected=%p, got=%p", thirdNode, third)
	}

	if fourth != fourthNode {
		t.Errorf("Expected=%p, got=%p", fourthNode, fourth)
	}

}

func TestLength(t *testing.T) {

	t.Run("with 3 nodes should have length 3", func(t *testing.T) {
		l := singly.NewLinkedList(1)
		l.Add(2)
		l.Add(3)

		want := 3
		got := l.Length()

		if want != got {
			t.Errorf("Expected=%d, but got=%d", want, got)
		}
	})

	t.Run("with only head should have length 1", func(t *testing.T) {
		l := singly.NewLinkedList(4)

		want := 1
		got := l.Length()

		if want != got {
			t.Errorf("Expected=%d, but got=%d", want, got)
		}
	})
}

func TestDelete(t *testing.T) {
	t.Run("should delete the head element", func(t *testing.T) {
		l := singly.NewLinkedList(4)
		l.Add(5)
		l.Add(6)

		expectedDeleteV := 4
		expectedLength := 2

		gotDel := l.Delete(0)
		gotLength := l.Length()

		if expectedDeleteV != gotDel.Data() {
			t.Errorf("Expected=%d, but got=%d", expectedDeleteV, gotDel.Data())
		}

		if expectedLength != gotLength {
			t.Errorf("Expected=%d, but got=%d", expectedLength, gotLength)
		}
	})

	t.Run("should delete the indexed element", func(t *testing.T) {
		l := singly.NewLinkedList(4)
		l.Add(5)
		l.Add(6)

		expectedDeleteV := 5
		expectedLength := 2

		gotDel := l.Delete(1)
		gotLength := l.Length()

		if expectedDeleteV != gotDel.Data() {
			t.Errorf("Expected=%d, but got=%d", expectedDeleteV, gotDel.Data())
		}

		if expectedLength != gotLength {
			t.Errorf("Expected=%d, but got=%d", expectedLength, gotLength)
		}
	})

	t.Run("should delete the last element", func(t *testing.T) {
		l := singly.NewLinkedList(4)
		l.Add(5)
		l.Add(6)

		expectedDeleteV := 6
		expectedLength := 2

		gotDel := l.Delete(2)
		gotLength := l.Length()

		if expectedDeleteV != gotDel.Data() {
			t.Errorf("Expected=%d, but got=%d", expectedDeleteV, gotDel.Data())
		}

		if expectedLength != gotLength {
			t.Errorf("Expected=%d, but got=%d", expectedLength, gotLength)
		}
	})

	t.Run("should panic if index is out of bounds", func(t *testing.T) {
		l := singly.NewLinkedList(4)

		defer func() {
			if r := recover(); r == nil {
				t.Errorf("expected func to panic, but it didn't")
			}
		}()

		l.Delete(-1)
	})
}
