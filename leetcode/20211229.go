package leetcode

import (
	"math"
)

var robedMapping map[*TreeNode]int

func robRe(root *TreeNode) int {

	if root == nil {
		return 0
	}

	if _, ok := robedMapping[root]; ok {
		return robedMapping[root]
	}

	// 盗窃这家
	sumA := root.Val
	if root.Left != nil {
		sumA += robRe(root.Left.Left) + robRe(root.Left.Right)
	}
	if root.Right != nil {
		sumA += robRe(root.Right.Left) + robRe(root.Right.Right)
	}
	// 不盗窃这家
	sumB := robRe(root.Left) + robRe(root.Right)
	if sumA > sumB {
		robedMapping[root] = sumA
		return sumA
	} else {
		robedMapping[root] = sumB
		return sumB
	}
}

func rob(root *TreeNode) int {
	robedMapping = make(map[*TreeNode]int)
	return robRe(root)
}

// 数组中是否存在和等于sum的元素组合
func canPartition(nums []int) bool {

	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	if sum%2 != 0 {
		return false
	}

	sum /= 2

	M, N := len(nums), sum
	dp := make([][]bool, 0, M)
	for i := 0; i < M; i++ {
		dp = append(dp, make([]bool, N+1))
	}

	for i := 0; i < M; i++ {
		dp[i][0] = true
	}

	for j := 1; j < N; j++ {
		if j == nums[0] {
			dp[0][j] = true
		}
	}

	for i := 0; i < M-1; i++ {
		for j := 0; j <= N; j++ {
			var flag bool
			if j-nums[i+1] >= 0 {
				flag = dp[i][j] || dp[i][j-nums[i+1]]
			} else {
				flag = dp[i][j]
			}
			dp[i+1][j] = flag
		}
	}
	return dp[M][N]
}

// 数组中和等于sum的元素组合的最小数量
func minimalCount(nums []int, sum int) int {

	M, N := len(nums), sum
	dp := make([][]int, 0, M)
	for i := 0; i < M; i++ {
		arr := make([]int, 0, N+1)
		for j := 0; j <= N; j++ {
			arr = append(arr, -1)
		}
		dp = append(dp, arr)
		dp[i][0] = 0
	}

	for j := 1; j <= N; j++ {
		if j == nums[0] {
			dp[0][j] = 1
		}
	}

	for i := 0; i < M-1; i++ {
		for j := 0; j <= N; j++ {
			if j-nums[i+1] >= 0 {
				count := math.MaxInt32
				if dp[i][j] >= 0 {
					count = dp[i][j]
				}
				if dp[i][j-nums[i+1]] >= 0 && dp[i][j-nums[i+1]]+1 < count {
					count = dp[i][j-nums[i+1]] + 1
				}
				if count == math.MaxInt32 {
					count = -1
				}
				dp[i+1][j] = count
			} else {
				dp[i+1][j] = dp[i][j]
			}
		}
	}

	return dp[M-1][N]
}
