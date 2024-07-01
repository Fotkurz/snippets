package binarysearch

import "cmp"

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
