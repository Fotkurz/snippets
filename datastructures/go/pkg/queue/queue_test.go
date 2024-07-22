package queue_test

import (
	"testing"

	"github.com/Fotkurz/snippets/datastructures/go/pkg/queue"
)

func TestIsEmpty(t *testing.T) {
	tests := []struct {
		desc string
		q    queue.Queue[int32]
		exp  bool
	}{
		{
			"return true",
			queue.Queue[int32]{},
			true,
		},
		{
			"return false",
			queue.Queue[int32]{1},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			value := tt.q.IsEmpty()
			if value != tt.exp {
				t.Errorf("expected=(%v), got=(%v)", tt.exp, value)
			}
		})
	}
}

func TestEnqueue(t *testing.T) {
	tests := []struct {
		desc      string
		q         queue.Queue[int32]
		push      int32
		expLength int
	}{
		{
			"length is 5 after push",
			queue.Queue[int32]{1, 2, 3, 4},
			5,
			5,
		},
		{
			"length is 1 if empty",
			queue.Queue[int32]{},
			1,
			1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			tt.q.Enqueue(tt.push)
			if tt.q.IsEmpty() {
				t.Error("expected queue to not be empty")
			}
			got := tt.q.Size()
			if got != tt.expLength {
				t.Errorf("expected size=(%d), got=(%d)", tt.expLength, got)
			}
		})
	}
}

func TestSize(t *testing.T) {
	t.Run("0 when len is 0", func(t *testing.T) {
		q := queue.Queue[int32]{}
		want := 0
		got := q.Size()

		if got != want {
			t.Errorf("expected size=(%d), got=(%d)", want, got)
		}
	})

	t.Run("2 when size is 2", func(t *testing.T) {
		q := queue.Queue[int32]{1, 2}
		want := 2
		got := q.Size()

		if got != want {
			t.Errorf("expected size=(%d), got=(%d)", want, got)
		}
	})
}

func TestDequeue(t *testing.T) {
	tests := []struct {
		desc    string
		q       queue.Queue[string]
		exp     string
		expSize int
	}{
		{
			"return last element of queue",
			queue.Queue[string]{"a", "b", "c"},
			"c",
			2,
		},
		{
			"return empty value when no elements",
			queue.Queue[string]{},
			"",
			0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			gotValue := tt.q.Dequeue()
			gotLength := tt.q.Size()

			if gotValue != tt.exp {
				t.Errorf("expected=(%s), got=(%s)", tt.exp, gotValue)
			}
			if gotLength != tt.expSize {
				t.Errorf("expected=(%d), got=(%d)", tt.expSize, gotLength)
			}
		})
	}
}

func TestPeek(t *testing.T) {
	tests := []struct {
		desc string
		q    queue.Queue[int8]
		exp  *int8
	}{
		{
			"return first element of queue",
			queue.Queue[int8]{1, 2},
			func() *int8 {
				v := int8(1)
				return &v
			}(),
		},
		{
			"return nil if queue is empty",
			queue.Queue[int8]{},
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			got := tt.q.Peek()

			if tt.exp == nil {
				if got != nil {
					t.Errorf("expected=(nil), got=(%v)", got)
				}
			} else if *got != *tt.exp {
				t.Errorf("expected=(%d), got=(%d)", *tt.exp, *got)
			}
		})
	}
}

func TestBack(t *testing.T) {
	tests := []struct {
		desc string
		q    queue.Queue[int8]
		exp  *int8
	}{
		{
			"return last element of queue",
			queue.Queue[int8]{1, 2},
			func() *int8 {
				v := int8(2)
				return &v
			}(),
		},
		{
			"return nil if queue is empty",
			queue.Queue[int8]{},
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			got := tt.q.Back()

			if tt.exp == nil {
				if got != nil {
					t.Errorf("expected=(nil), got=(%v)", got)
				}
			} else if *got != *tt.exp {
				t.Errorf("expected=(%d), got=(%d)", *tt.exp, *got)
			}
		})
	}
}
