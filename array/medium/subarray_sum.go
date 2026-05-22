package medium

func SubarraySum(nums []int, k int) int {
	mp := make(map[int]int, len(nums)+1)
	sum, res := 0, 0

	mp[0] = 1
	for i := range nums {
		sum += nums[i]
		res += mp[sum-k]
		mp[sum]++
	}
	return res
}
