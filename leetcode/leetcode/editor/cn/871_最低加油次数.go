package leetcode

//leetcode submit region begin(Prohibit modification and deletion)
func minRefuelStops(target int, startFuel int, stations [][]int) int {
	dp, n := make([]int, len(stations)+1), len(stations)
	dp[0] = startFuel
	for i := 0; i < n; i++ {
		for j := i; j >= 0; j-- {
			if dp[j] >= stations[i][0] {
				if dp[j] + stations[i][1] > dp[j+1] {
					dp[j+1] = dp[j] + stations[i][1]
				}
			}
		}
	}

	for i := 0; i <= n; i++ {
		if dp[i] >= target {
			return i
		}
	}

	return -1
}

//leetcode submit region end(Prohibit modification and deletion)
