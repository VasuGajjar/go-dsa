package answer

func FindKthMissing(arr []int, k int) int {
	n := len(arr)
	if n == 0 || arr[0] > k {
		return k
	}

	l, h := 0, n-1
	res := 0

	for l <= h {
		m := (l + h) / 2
		if v := arr[m] - m - 1; v < k {
			res = arr[m] + k - v
			l = m + 1
		} else {
			h = m - 1
		}
	}

	return res
}
