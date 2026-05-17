package array

func Reverse(nums []int) {
	i, j := 0, len(nums)-1

	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
