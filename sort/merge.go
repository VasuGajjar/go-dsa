package sort

func Merge(nums []int) {
	merge(nums, 0, len(nums)-1)
}

func merge(nums []int, start, end int) {
	if start >= end {
		return
	}

	mid := (start + end) / 2
	merge(nums, start, mid)
	merge(nums, mid+1, end)

	n := end - start + 1
	arr := make([]int, 0, n)

	i, j := start, mid+1
	for i <= mid && j <= end {
		if nums[i] < nums[j] {
			arr = append(arr, nums[i])
			i++
		} else {
			arr = append(arr, nums[j])
			j++
		}
	}

	for ; i <= mid; i++ {
		arr = append(arr, nums[i])
	}

	for ; j <= end; j++ {
		arr = append(arr, nums[j])
	}

	for i := range arr {
		nums[start+i] = arr[i]
	}
}
