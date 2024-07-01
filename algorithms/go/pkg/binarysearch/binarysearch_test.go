package binarysearch_test

import (
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
