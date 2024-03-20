package linkedlist_test

import (
	"fmt"
	"testing"

	"github.com/Fotkurz/datastructures/go/stack/pkg/linkedlist"
)

func TestAdd(t *testing.T) {
	headData := 12
	l := linkedlist.NewLinkedList(headData)

	l.Add(16)
	thirdNode := l.Add(18)

	headNode := l.Head()
	last := headNode.Next().Next()

	fmt.Println(&last)
	fmt.Println(&thirdNode)

	if headNode.Data() != headData {
		t.Errorf("Expected=%d, got=%d", headData, headNode.Data())
	}

	if thirdNode.Data() != last.Data() {
		t.Errorf("Expected=%d, got=%d", last.Data(), thirdNode.Data())
	}

	if &thirdNode != &last {
		t.Errorf("Expected=%d, got=%d", &last, &thirdNode)
	}

}
