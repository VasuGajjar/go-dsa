package easy

func LeftRotate(nums []int, k int) {
	n := len(nums)
	k = k % n

	reverse(nums[:k])
	reverse(nums[k:])
	reverse(nums)
}

func reverse(nums []int) {
	i, j := 0, len(nums)-1

	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
