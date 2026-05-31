package arrays

func LowerBound(nums []int, k int) int {
	low, high := 0, len(nums)-1
	var mid int

	for low <= high {
		mid = (low + high) / 2
		if k <= nums[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return low
}
