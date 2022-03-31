package leetcode

import "math"

//leetcode submit region begin(Prohibit modification and deletion)

var visitedMapping_path [][]bool

func minimumEffortPathRe(heights [][]int, i, j, pre, target int) bool {

	if i < 0 || i >= len(heights) || j < 0 || j >= len(heights[0]) ||
		visitedMapping_path[i][j] ||
		int(math.Abs(float64(heights[i][j]-pre))) > target {
		return false
	}

	if i == len(heights)-1 && j == len(heights[0])-1 {
		return true
	}

	visitedMapping_path[i][j] = true
	return minimumEffortPathRe(heights, i+1, j, heights[i][j], target) ||
		minimumEffortPathRe(heights, i-1, j, heights[i][j], target) ||
		minimumEffortPathRe(heights, i, j+1, heights[i][j], target) ||
		minimumEffortPathRe(heights, i, j-1, heights[i][j], target)
}

func minimumEffortPath(heights [][]int) int {

	m, n, l, r := len(heights), len(heights[0]), 0, 1000000
	for l < r {
		visitedMapping_path = make([][]bool, 0, m)
		for i := 0; i < m; i++ {
			visitedMapping_path = append(visitedMapping_path, make([]bool, n))
		}
		mid := (l + r) >> 1
		if minimumEffortPathRe(heights, 0, 0, heights[0][0], mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

//leetcode submit region end(Prohibit modification and deletion)
