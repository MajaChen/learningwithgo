package leetcode

import "math"

//leetcode submit region begin(Prohibit modification and deletion)

const INFINITY = 1000000

func findShortestWay(n, k, distanceThreshold int, matrix [][]int) []int {

	dp, s, res := make([]int, n), make(map[int]bool), make([]int, 0)
	s[k] = true
	for j := 0; j < n; j++ {
		dp[j] = matrix[k][j]
	}

	for len(s) < n {
		minVal, minIndex := math.MaxInt32, -1
		for i := 0; i < n; i++ {
			if !s[i] && dp[i] < minVal {
				minVal = dp[i]
				minIndex = i
			}
		}
		s[minIndex] = true

		for i := 0; i < n; i++ {
			if !s[i] && dp[minIndex]+matrix[minIndex][i] < dp[i] {
				dp[i] = dp[minIndex] + matrix[minIndex][i]
			}
		}

		if dp[minIndex] <= distanceThreshold {
			res = append(res, minIndex)
		} else {
			return res
		}
	}

	return res
}

func findTheCity(n int, edges [][]int, distanceThreshold int) int {

	matrix := make([][]int, 0, n)
	for i := 0; i < n; i++ {
		s := make([]int, 0, n)
		for j := 0; j < n; j++ {
			s = append(s, INFINITY)
		}
		matrix = append(matrix, s)
	}

	for _, edge := range edges {
		matrix[edge[0]][edge[1]] = edge[2]
		matrix[edge[1]][edge[0]] = edge[2]
	}

	minVal, minIndex := math.MaxInt32, -1
	for i := 0; i < n; i++ {
		if j := len(findShortestWay(n, i, distanceThreshold, matrix)); j <= minVal {
			minVal = j
			minIndex = i
		}
	}

	return minIndex
}

//leetcode submit region end(Prohibit modification and deletion)
