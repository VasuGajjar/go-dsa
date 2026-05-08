package easy

func SingleNumberWithMap(nums []int) int {
	n := len(nums)
	mp := make(map[int]bool, (n/2)+1)

	for i := range n {
		if mp[nums[i]] {
			delete(mp, nums[i])
		} else {
			mp[nums[i]] = true
		}
	}

	for k := range mp {
		return k
	}

	return -1
}

func SingleNumberWithXor(nums []int) int {
	res := 0

	for i := range nums {
		res ^= nums[i]
	}

	return res
}
