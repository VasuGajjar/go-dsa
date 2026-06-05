package arrays

func SingleNonDuplicate(nums []int) int {
	low, high := 0, len(nums)-1
	mid := 0

	for low <= high {
		mid = (low + high) / 2

		if mid > 0 && nums[mid] == nums[mid-1] {
			if (mid-1)%2 == 0 {
				low = mid + 1
			} else {
				high = mid - 1
			}
		} else if mid < high && nums[mid] == nums[mid+1] {
			if mid%2 == 0 {
				low = mid + 1
			} else {
				high = mid - 1
			}
		} else {
			return nums[mid]
		}
	}

	return nums[low]
}
