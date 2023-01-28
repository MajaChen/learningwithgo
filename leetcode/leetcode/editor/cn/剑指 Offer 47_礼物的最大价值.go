package leetcode

//leetcode submit region begin(Prohibit modification and deletion)

func getValue(i, j, m, n int, grid, dp [][]int) int {
	r, d := grid[i][j], grid[i][j]
	// 向右移动
	if j < n-1 {
		r += dp[i][j+1]
	}
	// 向下移动
	if i < m-1 {
		d += dp[i+1][j]
	}
	if r < d {
		r = d
	}
	return r
}

func maxValue(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, 0, m)
	for i := 0; i < m; i++ {
		dp = append(dp, make([]int, n))
	}
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			dp[i][j] = getValue(i, j, m, n, grid, dp)
		}
	}
	return dp[0][0]
}

//leetcode submit region end(Prohibit modification and deletion)
