package arrays

func SearchRange(nums []int, k int) []int {
	n := len(nums) - 1
	lower := LowerBound(nums, k)

	if lower < 0 || lower > n || nums[lower] != k {
		return []int{-1, -1}
	}

	upper := UpperBound(nums, k)
	return []int{lower, upper - 1}
}
