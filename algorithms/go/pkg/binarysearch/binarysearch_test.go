package binarysearch_test

import (
	"errors"
	"slices"
	"testing"

	"github.com/Fotkurz/algorithms/go/pkg/binarysearch"
	"github.com/Fotkurz/algorithms/go/pkg/testutils"
)

func TestBinarySearch(t *testing.T) {

	t.Run("should return the index of the value when exist", func(t *testing.T) {
		var want int = 3
		got := binarysearch.BinarySearch[[]int32]([]int32{1, 2, 3, 4, 5, 6}, 4)
		testutils.AssertEqual(t, want, got)
	})

	t.Run("should return -1 when value is not found", func(t *testing.T) {
		var want int = -1
		got := binarysearch.BinarySearch[[]string]([]string{"a", "b", "c", "d", "e"}, "z")
		testutils.AssertEqual(t, want, got)
	})

	t.Run("test with the slices lib binarysearch", func(t *testing.T) {
		var want int = 3
		got, ok := slices.BinarySearch([]int32{1, 2, 3, 4, 5, 6}, 4)

		testutils.AssertEqual(t, want, got)
		testutils.AssertEqual(t, true, ok)
	})
}

func TestBinarySearchInt(t *testing.T) {
	tests := []struct {
		name    string
		slc     []int
		value   int
		want    int
		wantErr error
	}{
		{
			name:    "return -1 and err if slice is empty",
			slc:     []int{},
			value:   1,
			want:    -1,
			wantErr: errors.New("slice is empty"),
		},
		{
			name:    "return -1 and err if value not in slice",
			slc:     []int{10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36},
			value:   120,
			want:    -1,
			wantErr: errors.New("120 not in slice"),
		},
		{
			name:    "return 3",
			slc:     []int{10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36},
			value:   16,
			want:    3,
			wantErr: nil,
		},
		{
			name:    "return 9",
			slc:     []int{10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36},
			value:   28,
			want:    9,
			wantErr: nil,
		},
		{
			name:    "return 20",
			slc:     []int{10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36},
			value:   20,
			want:    5,
			wantErr: nil,
		},
		{
			name:    "return mid",
			slc:     []int{10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36},
			value:   22,
			want:    6,
			wantErr: nil,
		},
		{
			name:    "return first",
			slc:     []int{10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36},
			value:   10,
			want:    0,
			wantErr: nil,
		},
		{
			name:    "return last",
			slc:     []int{10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36},
			value:   36,
			want:    13,
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := binarysearch.SafeBinarySearchInt(tt.slc, tt.value)
			if tt.wantErr != nil {
				testutils.AssertEqual(t, tt.wantErr.Error(), gotErr.Error())
			} else {
				testutils.AssertNil(t, gotErr)
			}

			testutils.AssertEqual(t, tt.want, got)
		})
	}
}

func BenchmarkBinarySearchPerformance(b *testing.B) {
	//  go test -bench=. -cpu=1 ./pkg/binarysearch/...
	limit := 1000000
	var slc []int
	for i := 0; i < limit; i++ {
		slc = append(slc, i+2)
	}

	b.Run("benchmark BinarySearch", func(b *testing.B) {
		binarysearch.BinarySearch[[]int](slc, 2)
	})

	b.Run("benchmark SafeBinarySearchInt", func(b *testing.B) {
		binarysearch.SafeBinarySearchInt(slc, 2)
	})
}
