package linkedlist_test

import (
	"testing"

	"github.com/Fotkurz/datastructures/go/stack/pkg/linkedlist"
)

func TestAdd(t *testing.T) {
	headData := 12
	l := linkedlist.NewLinkedList(headData)

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
