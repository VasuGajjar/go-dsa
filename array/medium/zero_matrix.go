package medium

func ZeroMatrix(matrix [][]int) {
	col0 := false

	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] != 0 {
				continue
			}

			matrix[i][0] = 0
			if j == 0 {
				col0 = true
			} else {
				matrix[0][j] = 0
			}
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if matrix[0][0] == 0 {
		for j := range matrix[0] {
			matrix[0][j] = 0
		}
	}

	if col0 {
		for i := range matrix {
			matrix[i][0] = 0
		}
	}
}
