package hard

func PascalTriangle(n int) [][]int {
	res := make([][]int, n)

	for k := range n {
		res[k] = make([]int, k+1)

		for i := range res[k] {
			if i == 0 || i == k {
				res[k][i] = 1
				continue
			}

			res[k][i] = res[k-1][i-1] + res[k-1][i]
		}
	}

	return res
}
