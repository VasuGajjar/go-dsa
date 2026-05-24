package medium

func SpiralOrder(matrix [][]int) []int {
	n, m := len(matrix), len(matrix[0])
	res := make([]int, 0, n*m)

	t, b, l, r := 0, n-1, 0, m-1

	for t <= b && l <= r {
		for i := l; i <= r; i++ {
			res = append(res, matrix[t][i])
		}
		t++

		if l <= r {
			for i := t; i <= b; i++ {
				res = append(res, matrix[i][r])
			}
			r--
		}

		if t <= b {
			for i := r; i >= l; i-- {
				res = append(res, matrix[b][i])
			}
			b--
		}

		if l <= r {
			for i := b; i >= t; i-- {
				res = append(res, matrix[i][l])
			}
			l++
		}
	}

	return res
}
