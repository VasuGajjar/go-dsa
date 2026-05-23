package arrays

func BinarySearch(nums []int, val int) int {
	l, h := 0, len(nums)-1

	for l <= h {
		mid := (l + h) / 2
		if nums[mid] == val {
			return mid
		} else if nums[mid] < val {
			l = mid + 1
		} else {
			h = mid - 1
		}
	}

	return -1
}
