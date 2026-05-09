package easy

func SubarraySum(nums []int, k int) int {
	res, sum := 0, 0
	mp := make(map[int]int, len(nums))
	mp[0] = -1

	for i := range nums {
		sum += nums[i]

		if v, ok := mp[sum-k]; ok {
			res = max(res, i-v)
		}

		if _, ok := mp[sum]; !ok {
			mp[sum] = i
		}
	}

	return res
}
