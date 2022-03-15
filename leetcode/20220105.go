package leetcode

//"abcde"
//"ace"
func longestCommonSubsequence(text1 string, text2 string) int {

	m, n := len(text1), len(text2)
	dp := make([][]int, 0, m+1)
	for i := 0; i < m+1; i++ {
		dp = append(dp, make([]int, n+1))
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				if dp[i][j] = dp[i-1][j]; dp[i][j] < dp[i][j-1] {
					dp[i][j] = dp[i][j-1]
				}
			}
		}
	}

	return dp[m][n]
}

func findLength(nums1 []int, nums2 []int) int {

	m, n, maximalLength := len(nums1), len(nums2), 0
	dp := make([][]int, 0, m+1)
	for i := 0; i < m+1; i++ {
		dp = append(dp, make([]int, n+1))
	}

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if nums1[i] == nums2[j] {
				if dp[i][j] = dp[i+1][j+1] + 1; dp[i][j] > maximalLength {
					maximalLength = dp[i][j]
				}
			} else {
				dp[i][j] = 0
			}
		}
	}

	return maximalLength
}

func findLength2(nums1 []int, nums2 []int) int {

	m, n, maximalLength := len(nums1), len(nums2), 0
	dp := make([][]int, 0, m+1)
	for i := 0; i < m+1; i++ {
		dp = append(dp, make([]int, n+1))
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if nums1[i-1] == nums2[j-1] {
				if dp[i][j] = dp[i-1][j-1] + 1; dp[i][j] > maximalLength {
					maximalLength = dp[i][j]
				}
			} else {
				dp[i][j] = 0
			}
		}
	}
	return maximalLength
}

func lengthOfLIS(nums []int) int {

	dp, maximalLength := make([]int, len(nums)), 0
	for i := 0; i < len(nums); i++ {
		// dp[i]
		var k int
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] && dp[j] > k {
				k = dp[j]
			}
		}
		if dp[i] = k + 1; dp[i] > maximalLength {
			maximalLength = dp[i]
		}
	}
	return maximalLength
}
