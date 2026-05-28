package hard

import "math"

func MajorityElement(nums []int) []int {
	res := make([]int, 0, 2)

	v1, v2, c1, c2 := math.MinInt, math.MinInt, 0, 0
	for i := range nums {
		if nums[i] == v1 {
			c1++
		} else if nums[i] == v2 {
			c2++
		} else if c1 == 0 {
			v1 = nums[i]
			c1 = 1
		} else if c2 == 0 {
			v2 = nums[i]
			c2 = 1
		} else {
			c1--
			c2--
		}
	}

	c1, c2 = 0, 0
	for i := range nums {
		if nums[i] == v1 {
			c1++
		} else if nums[i] == v2 {
			c2++
		}
	}

	n := len(nums) / 3
	if c1 > n {
		res = append(res, v1)
	}
	if c2 > n {
		res = append(res, v2)
	}

	return res
}
