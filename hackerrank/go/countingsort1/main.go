package main

// CountingSort
//
//	Returns a frequency of values for arr using the
//	largest number as ref
func CountingSort(arr []int32, max int32) []int32 {
	result := make([]int32, max)

	for i := 0; i < len(arr); i++ {
		if i >= 0 {
			result[arr[i]]++
		}
	}

	return result
}
