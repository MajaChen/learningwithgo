package leetcode

//leetcode submit region begin(Prohibit modification and deletion)
func cuttingRope(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i <= n; i++ {
		ml := 0
		for j := 1; j < i; j++ {
			l := j
			r := dp[i-j]
			if r < i-j {
				r = i - j
			}
			if l*r > ml {
				ml = l * r
			}
		}
		dp[i] = ml
	}
	return dp[n]
}

//leetcode submit region end(Prohibit modification and deletion)
