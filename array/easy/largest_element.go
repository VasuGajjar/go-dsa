package easy

import "math"

func LargestElement(nums []int) int {
	mx := math.MinInt

	for _, i := range nums {
		mx = max(mx, i)
	}

	return mx
}

func SecondLargestElement(nums []int) int {
	mx, secondMx := math.MinInt, math.MinInt

	for i := 0; i < len(nums); i++ {
		if nums[i] > mx {
			secondMx = mx
			mx = nums[i]
		} else if nums[i] > secondMx {
			secondMx = nums[i]
		}
	}

	return secondMx
}
