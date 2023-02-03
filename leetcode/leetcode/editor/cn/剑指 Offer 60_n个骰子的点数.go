package leetcode

//leetcode submit region begin(Prohibit modification and deletion)
func dicesProbability(n int) []float64 {
	dp := make([]float64, 7)
	for i := 1; i <= 6; i++ {
		dp[i] = float64(1) / float64(6)
	}
	for i := 2; i <= n; i++ {
		tmp := make([]float64, 6*i+1)
		for j := i - 1; j <= 6*(i-1); j++ {
			for k := 1; k <= 6; k++ {
				tmp[j+k] += dp[j] / float64(6)
			}
		}
		dp = tmp
	}
	return dp[n:]
}

//leetcode submit region end(Prohibit modification and deletion)
