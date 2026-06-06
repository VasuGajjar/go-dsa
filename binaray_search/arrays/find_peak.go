package arrays

func FindPeak(nums []int) int {
	n := len(nums)
	low, high := 0, n-1
	var mid int

	for low <= high {
		mid = (low + high) / 2
		if ml, mr := mid-1, mid+1; (ml < 0 || nums[ml] < nums[mid]) && (mr >= n || nums[mr] < nums[mid]) {
			return mid
		} else if ml >= 0 && nums[ml] > nums[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}
