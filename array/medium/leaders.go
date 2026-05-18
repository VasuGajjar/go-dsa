package medium

import (
	"fmt"
)

func Leaders(nums []int) string {
	n := len(nums)
	mx := nums[n-1]
	res := fmt.Sprintf("%v", mx)

	for i := n - 2; i >= 0; i-- {
		if nums[i] > mx {
			res = fmt.Sprintf("%v %v", nums[i], res)
			mx = nums[i]
		}
	}

	return res
}
