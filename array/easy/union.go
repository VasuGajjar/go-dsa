package easy

func Union(nums1, nums2 []int) []int {
	n, m := len(nums1), len(nums2)
	res := make([]int, 0, n+m)

	i, j := 0, 0

	for i < n && j < m {
		if nums1[i] < nums2[j] {
			if k := len(res); k <= 0 || res[k-1] != nums1[i] {
				res = append(res, nums1[i])
			}
			i++
		} else {
			if k := len(res); k <= 0 || res[k-1] != nums2[j] {
				res = append(res, nums2[j])
			}
			j++
		}
	}

	for ; i < n; i++ {
		if k := len(res); k <= 0 || res[k-1] != nums1[i] {
			res = append(res, nums1[i])
		}
	}

	for ; j < m; j++ {
		if k := len(res); k <= 0 || res[k-1] != nums2[j] {
			res = append(res, nums2[j])
		}
	}

	return res
}
