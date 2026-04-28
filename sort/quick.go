package sort

func Quick(nums []int) {
	if len(nums) <= 1 {
		return
	}

	p := partition(nums)
	Quick(nums[:p])
	Quick(nums[p+1:])
}

func partition(nums []int) int {
	p, n := 0, len(nums)-1
	i, j := p+1, n

	if n <= 0 {
		return 0
	}

	for i <= n && j > p {
		for ; i <= n && nums[i] <= nums[p]; i++ {
		}

		for ; j > p && nums[j] > nums[p]; j-- {
		}

		if i >= j {
			break
		}

		nums[i], nums[j] = nums[j], nums[i]
		i++
	}

	nums[p], nums[j] = nums[j], nums[p]
	return j
}
