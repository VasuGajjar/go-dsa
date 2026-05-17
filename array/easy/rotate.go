package easy

import "github.com/VasuGajjar/go-dsa/array"

func LeftRotate(nums []int, k int) {
	n := len(nums)
	k = k % n

	array.Reverse(nums[:k])
	array.Reverse(nums[k:])
	array.Reverse(nums)
}
