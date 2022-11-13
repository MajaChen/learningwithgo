package leetcode

//leetcode submit region begin(Prohibit modification and deletion)
func numDistinct(s string, t string) int {
	m, n := len(s), len(t)
	dp := make([][]int, 0, m+1)
	for i := 0; i <= m; i++ {
		dp = append(dp, make([]int, n+1))
	}
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if i == 0 {
				dp[i][j] = 0
			}
			if j == 0 {
				dp[i][j] = 1
			}
		}
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = dp[i-1][j]
			if s[i-1] == t[j-1] {
				dp[i][j] += dp[i-1][j-1]
			}
		}
	}

	return dp[m][n]
}

//leetcode submit region end(Prohibit modification and deletion)
