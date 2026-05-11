package medium

func SortColorsWithCount(nums []int) {
	cnt0, cnt1 := 0, 0

	for _, i := range nums {
		if i == 0 {
			cnt0++
		} else if i == 1 {
			cnt1++
		}
	}

	i := 0

	for range cnt0 {
		nums[i] = 0
		i++
	}

	for range cnt1 {
		nums[i] = 1
		i++
	}

	for ; i < len(nums); i++ {
		nums[i] = 2
	}
}

func SortColorsWithDutchFlagAlgorithm(nums []int) {
	i, j := 0, len(nums)

	for k := i; k < j; {
		switch nums[k] {
		case 0:
			nums[k], nums[i] = nums[i], nums[k]
			i++
			k++
		case 1:
			k++
		case 2:
			nums[k], nums[j-1] = nums[j-1], nums[k]
			j--
		}
	}
}
