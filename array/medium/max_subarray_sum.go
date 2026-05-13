package medium

func MaxSubarraySum(nums []int) int {
	sum, mx := 0, 0

	for i := range nums {
		sum += nums[i]
		mx = max(mx, sum)

		if sum < 0 {
			sum = 0
		}
	}

	return mx
}
