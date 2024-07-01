package selectionsort

func SelectionSortFunc[T any](slc []T, f func(t []T) T) []T {
	arr := []T{}

	for i := 0; i < len(slc); i++ {
		value := f(slc)
		arr = append(arr, value)
	}

	return arr
}
