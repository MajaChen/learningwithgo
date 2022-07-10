package leetcode

//leetcode submit region begin(Prohibit modification and deletion)
func lenLongestFibSubseq(arr []int) int {
	mapping := make(map[int]int)
	for i, a := range arr {
		mapping[a] = i
	}

	n, longestLen := len(arr), 0
	dp := make([][]int, 0, n)
	for i := 0; i < n; i++ {
		dp = append(dp, make([]int, n))
	}

	for i := 1; i < n; i++ {
		for j := i + 1; j < n && arr[j]-arr[i] < arr[i]; j++ {
			if k, ok := mapping[arr[j]-arr[i]]; ok && k < i {
				if dp[k][i] > 0 {
					dp[i][j] = dp[k][i] + 1
				} else {
					dp[i][j] = 3
				}
				if dp[i][j] > longestLen {
					longestLen = dp[i][j]
				}
			}
		}
	}

	return longestLen
}

//leetcode submit region end(Prohibit modification and deletion)
