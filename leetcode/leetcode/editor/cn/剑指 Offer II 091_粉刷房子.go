package leetcode

import "math"

//leetcode submit region begin(Prohibit modification and deletion)
func minCost(costs [][]int) int {
	n, m := len(costs), 3
	dp := make([][]int, 0, n)
	for i := 0; i < n; i++ {
		dp = append(dp, make([]int, m))
	}

	dp[0][0], dp[0][1], dp[0][2] = costs[0][0], costs[0][1], costs[0][2]
	for i := 1; i < n; i++ {
		dp[i][0] = costs[i][0] + int(math.Min(float64(dp[i-1][1]), float64(dp[i-1][2])))
		dp[i][1] = costs[i][1] + int(math.Min(float64(dp[i-1][0]), float64(dp[i-1][2])))
		dp[i][2] = costs[i][2] + int(math.Min(float64(dp[i-1][0]), float64(dp[i-1][1])))
	}

	return int(math.Min(math.Min(float64(dp[n-1][0]), float64(dp[n-1][1])), float64(dp[n-1][2])))
}

//leetcode submit region end(Prohibit modification and deletion)
