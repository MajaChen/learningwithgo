package leetcode

//leetcode submit region begin(Prohibit modification and deletion)
func numSpecial(mat [][]int) int {
	m, n := len(mat), len(mat[0])
	if m == 0 || n == 0 {
		return 0
	}
	row, col, ps := make([]int, m), make([]int, n), make([][2]int, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 1 {
				row[i]++
				col[j]++
				ps = append(ps, [2]int{i, j})
			}
		}
	}

	var ans int
	for _, p := range ps {
		if row[p[0]] == 1 && col[p[1]] == 1 {
			ans++
		}
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
