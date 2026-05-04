package easy

func MoveZeros(nums []int) {
	i := 0

	for _, v := range nums {
		if v != 0 {
			nums[i] = v
			i++
		}
	}

	for ; i < len(nums); i++ {
		nums[i] = 0
	}
}
