package medium

import "github.com/VasuGajjar/go-dsa/array"

func NextPermutation(nums []int) {
	n := len(nums) - 1
	i := n - 1

	for ; i >= 0 && nums[i] >= nums[i+1]; i-- {
	}

	array.Reverse(nums[i+1:])

	if i >= 0 {
		j := i + 1

		for ; j < n && nums[i] >= nums[j]; j++ {
		}

		nums[i], nums[j] = nums[j], nums[i]
	}
}
