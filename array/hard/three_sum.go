package hard

import "sort"

func ThreeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)

	for i := range len(nums) - 2 {
		if i != 0 && nums[i-1] == nums[i] {
			continue
		}

		j, k := i+1, len(nums)-1
		for j < k {
			if val := nums[i] + nums[j] + nums[k]; val == 0 {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				j++
				k--
				for j < k && nums[j] == nums[j-1] {
					j++
				}
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			} else if val < 0 {
				j++
			} else {
				k--
			}
		}
	}

	return res
}
