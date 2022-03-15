package leetcode

func setZeroes(matrix [][]int) {
	n, m := len(matrix), len(matrix[0])
	var rowZeros, colZeros bool
	for j := 0; j < m; j++ {
		if matrix[0][j] == 0 {
			rowZeros = true
			break
		}
	}

	for i := 0; i < n; i++ {
		if matrix[i][0] == 0 {
			colZeros = true
			break
		}
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for j := 1; j < m; j++ {
		if matrix[0][j] == 0 {
			for i := 1; i < n; i++ {
				matrix[i][j] = 0
			}
		}
	}

	for i := 1; i < n; i++ {
		if matrix[i][0] == 0 {
			for j := 1; j < m; j++ {
				matrix[i][j] = 0
			}
		}
	}

	if rowZeros {
		for j := 0; j < m; j++ {
			matrix[0][j] = 0
		}
	}

	if colZeros {
		for i := 0; i < n; i++ {
			matrix[i][0] = 0
		}
	}
	return
}
