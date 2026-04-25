package sort

func Bubble(nums []int) {
	n := len(nums)

	for j := range n - 1 {
		for i := 1; i < n-j; i++ {
			if nums[i-1] > nums[i] {
				nums[i-1], nums[i] = nums[i], nums[i-1]
			}
		}
	}
}
