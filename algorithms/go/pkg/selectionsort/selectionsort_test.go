package selectionsort_test

import (
	"math"
	"testing"

	"github.com/Fotkurz/algorithms/go/pkg/selectionsort"
	"github.com/Fotkurz/algorithms/go/pkg/testutils"
)

func TestSelectionSortFunc(t *testing.T) {
	t.Run("should return a slice in descending order", func(t *testing.T) {
		want := []int32{6, 5, 4, 2, 1}
		got := selectionsort.SelectionSortFunc([]int32{5, 2, 6, 1, 4}, func(t []int32) int32 {
			min := math.MaxInt32
			for _, v := range t {
				if int(v) < min {
					min = int(v)
				}
			}

			return int32(min)
		})

		testutils.AssertEqual(t, want, got)
	})
}
