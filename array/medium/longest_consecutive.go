package medium

import "sort"

func LongestConsecutiveWithMap(nums []int) int {
	mp := make(map[int]bool)
	for i := range nums {
		mp[nums[i]] = true
	}

	mx, cnt := 0, 0
	for k := range mp {
		if _, ok := mp[k-1]; ok {
			continue
		}
		cnt = 0
		for i := k; mp[i]; i++ {
			cnt++
		}
		mx = max(mx, cnt)
	}
	return mx
}

func LongestConsecutiveWithSort(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	sort.Ints(nums)
	mx, cnt := 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i-1] == nums[i] {
			continue
		} else if nums[i-1]+1 == nums[i] {
			cnt++
			mx = max(mx, cnt)
		} else {
			cnt = 1
		}
	}

	return mx
}
