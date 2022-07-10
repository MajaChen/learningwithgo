package leetcode

import "math"

//leetcode submit region begin(Prohibit modification and deletion)

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

func min(i, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}

func cherryPickup(grid [][]int) int {
	n := len(grid)
	dp := make([][][]int, 0, 2*n-1)
	for i := 0; i < 2*n-1; i++ {
		arr := make([][]int, 0, n)
		for j := 0; j < n; j++ {
			subArr := make([]int, 0, n)
			for k := 0; k < n; k++ {
				subArr = append(subArr, math.MinInt32)
			}
			arr = append(arr, subArr)
		}
		dp = append(dp, arr)
	}

	dp[0][0][0] = grid[0][0]
	for k := 1; k < 2*n-1; k++ {
		for x1 := max(0, k-n+1); x1 <= min(n-1, k); x1++ {
			y1 := k - x1
			if grid[x1][y1] == -1 {
				continue
			}
			for x2 := max(0, k-n+1); x2 <= min(n-1, k); x2++ {
				y2 := k - x2
				if grid[x2][y2] == -1 {
					continue
				}
				res := dp[k-1][x1][x2]
				if x1 > 0 {
					res = max(res, dp[k-1][x1-1][x2])
				}
				if x2 > 0 {
					res = max(res, dp[k-1][x1][x2-1])
				}
				if x1 > 0 && x2 > 0 {
					res = max(res, dp[k-1][x1-1][x2-1])
				}
				res += grid[x1][y1]
				if x1 != x2 {
					res += grid[x2][y2]
				}
				dp[k][x1][x2] = res
			}
		}
	}

	return max(0, dp[2*n-2][n-1][n-1])
}

//leetcode submit region end(Prohibit modification and deletion)
