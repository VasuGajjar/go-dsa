package medium

func MajorityElement(nums []int) int {
	n := len(nums)

	if n < 1 {
		return -1
	}

	val, cnt := nums[0], 1

	for i := 1; i < n; i++ {
		if nums[i] == val {
			cnt++
		} else {
			cnt--
		}

		if cnt <= 0 {
			val, cnt = nums[i], 1
		}
	}

	return val
}
