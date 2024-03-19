package linkedlist_test

import (
	"fmt"
	"testing"

	"github.com/Fotkurz/datastructures/go/stack/pkg/linkedlist"
)

func TestAdd(t *testing.T) {
	initial := 12
	l := linkedlist.NewLinkedList(initial)

	l.Add(16)
	res := l.Add(18)

	headNode := l.Head()
	last := headNode.Next().Next()

	fmt.Println(&last)
	fmt.Println(&res)

	if headNode.Data() != initial {
		t.Errorf("Expected=%d, got=%d", initial, headNode.Data())
	}

	if res.Data() != last.Data() {
		t.Errorf("Expected=%d, got=%d", last.Data(), res.Data())
	}

	if &res != &last {
		t.Errorf("Expected=%d, got=%d", &last, &res)
	}

}
