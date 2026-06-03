package arrays

import "math"

func FindMin(nums []int) int {
	low, high := 0, len(nums)-1
	ans := math.MaxInt

	for low <= high {
		mid := (low + high) / 2
		ans = min(ans, nums[mid])
		if nums[low] <= nums[mid] && nums[mid] <= nums[high] {
			ans = min(ans, nums[low])
			break
		} else if nums[low] <= nums[mid] && nums[mid] >= nums[high] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return ans
}
