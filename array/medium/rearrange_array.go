package medium

func RearrangeArray(nums []int) []int {
	res := make([]int, len(nums))
	j, k := 0, 1

	for i := range nums {
		if nums[i] >= 0 {
			res[j] = nums[i]
			j += 2
		} else {
			res[k] = nums[i]
			k += 2
		}
	}

	return res
}