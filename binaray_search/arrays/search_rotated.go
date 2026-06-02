package arrays

func SearchRotatedI(nums []int, k int) int {
	low, high := 0, len(nums)-1

	for low <= high {
		mid := (low + high) / 2
		if nums[mid] == k {
			return mid
		} else if nums[low] <= nums[mid] {
			if nums[low] <= k && k <= nums[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else {
			if nums[mid] <= k && k <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}

	return -1
}

func SearchRotatedII(nums []int, k int) bool {
	low, high := 0, len(nums)-1

	for low <= high {
		mid := (low + high) / 2
		if nums[mid] == k {
			return true
		} else if nums[low] == nums[mid] && nums[mid] == nums[high] {
			low, high = low+1, high-1
		} else if nums[low] < nums[mid] {
			if nums[low] <= k && k <= nums[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else {
			if nums[mid] <= k && k <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}

	return false
}
