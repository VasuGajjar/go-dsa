package easy

func RemoveDuplicates(nums []int) {
	i := 1

	for j := 1; j < len(nums); j++ {
		if nums[j-1] != nums[j] {
			nums[i] = nums[j]
			i++
		}
	}

	for ; i < len(nums); i++ {
		nums[i] = 0
	}
}
