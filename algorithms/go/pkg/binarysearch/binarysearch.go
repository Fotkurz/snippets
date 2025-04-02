package binarysearch

import (
	"cmp"
	"errors"
	"fmt"
)

// BinarySearch simple execute a binary search in a ordered slice.
// returns the index of the value or -1 if it is not found
func BinarySearch[S ~[]T, T cmp.Ordered](slc []T, value T) int {
	var low, index int = 0, 0
	high := len(slc) - 1

	for low <= high {
		mid := (low + high) / 2
		kick := slc[mid]

		if cmp.Compare(kick, value) == 0 {
			return index
		}
		if cmp.Compare(kick, value) == +1 {
			high = mid - 1
		} else {
			low = mid + 1
			index = low
		}
	}

	return -1
}

func SafeBinarySearchInt(slc []int, value int) (int, error) {
	if len(slc) == 0 {
		return -1, errors.New("slice is empty")
	}

	high := len(slc) - 1
	low := 0

	for low <= high {
		i := (high + low) / 2
		mid := slc[i]
		if mid == value {
			return i, nil
		}
		if mid > value {
			high = i - 1
		} else {
			low = i + 1
		}
	}

	return -1, fmt.Errorf("%v not in slice", value)
}
