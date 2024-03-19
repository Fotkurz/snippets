package stack_test

import (
	"testing"

	"github.com/Fotkurz/datastructures/go/stack/pkg/stack"
)

func TestPop(t *testing.T) {
	tests := []struct {
		name string
		init stack.Stack[string]
		exp1 stack.Stack[string]
		exp2 string
		exp3 bool
	}{
		{
			"should return 1,2,3",
			[]string{"1", "2", "3", "4"},
			[]string{"1", "2", "3"},
			"4",
			true,
		},
		{
			"should return empty",
			[]string{""},
			[]string{},
			"",
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			top, ok := tt.init.Pop()

			if len(tt.init) != len(tt.exp1) {
				t.Errorf("Expected len=%v, got len=%v", len(tt.exp1), len(tt.init))
			}

			if tt.exp2 != top {
				t.Errorf("Expected=%v, got=%v", tt.exp2, top)
			}

			if ok == false {
				t.Errorf("Expected=%v, got=%v", tt.exp3, ok)
			}

		})
	}
}

func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		init stack.Stack[string]
		arg1 string
		exp  stack.Stack[string]
	}{
		{
			"should add 4 to the last position",
			[]string{"1", "2", "3"},
			"4",
			[]string{"1", "2", "3", "4"},
		},
		{
			"should add 2 to the first position if s is empty",
			[]string{},
			"2",
			[]string{"2"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.init.Add(tt.arg1)

			if len(tt.init) != len(tt.exp) {
				t.Errorf("Expected=%v, got=%v", tt.exp, tt.init)
			}

			if tt.init[len(tt.init)-1] != tt.arg1 {
				t.Errorf("Expected=%v, got=%v", tt.arg1, tt.init[len(tt.init)-1])
			}
		})
	}
}

func TestPeek(t *testing.T) {
	tests := []struct {
		name string
		init stack.Stack[int32]
		exp  int32
	}{
		{
			"should return the value 3",
			[]int32{1, 2, 3},
			3,
		},
		{
			"should return empty if stack is empty",
			[]int32{},
			*new(int32),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := tt.init.Peek()

			if res != tt.exp {
				t.Errorf("Expected=%v, got=%v", tt.exp, res)
			}
		})
	}
}
