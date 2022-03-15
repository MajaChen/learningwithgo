package leetcode

func traverse(matrix [][]int, startLine, startCol, span int) bool {
	for i, j, k := startLine, startCol, 0; k < span; k++ {
		if matrix[i+k][j] != 0 || matrix[i][j+k] != 0 ||
			matrix[i+k][j+span-1] != 0 || matrix[i+span-1][j+k] != 0 {
			return false
		}
	}
	return true
}

func findSquare2(matrix [][]int) []int {

	if len(matrix) <= 0 {
		return nil
	}

	n, maxSpan, ai, aj := len(matrix), 1, -1, -1
	for span := 1; span <= n; span++ {
		bi, bj := -1, -1
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if i+span <= n && j+span <= n && traverse(matrix, i, j, span) {
					maxSpan = span
					if bi < 0 || bj < 0 {
						bi, bj = i, j
					}
				}
			}
		}
		if bi >= 0 || bj >= 0 {
			maxSpan, ai, aj = span, bi, bj
		}
	}

	if ai < 0 || aj < 0 {
		return []int{}
	}

	return []int{ai, aj, maxSpan}
}
