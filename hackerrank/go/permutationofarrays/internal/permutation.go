package internal

import "fmt"

func permutations(res, nums []int32, l, h int) {

	if l == h {
		res = append(res, nums...)
		for i := 0; i < len(nums); i++ {
			fmt.Printf("%d ", nums[i])
		}
		fmt.Print("\n")
		return
	}

	for i := l; i < h; i++ {
		// swapping
		nums[l], nums[i] = nums[i], nums[l]

		permutations(res, nums, l+1, h)

		nums[l], nums[i] = nums[i], nums[l]
	}
}

func Permute(nums []int32) []int32 {
	length := len(nums) - 1
	res := []int32{}

	permutations(res, nums, 0, length)
	return res
}
