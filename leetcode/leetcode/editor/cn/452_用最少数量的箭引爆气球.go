package leetcode

import "sort"

//leetcode submit region begin(Prohibit modification and deletion)
func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(x, y int) bool {
		return points[x][1] < points[y][1] ||
			(points[x][1] == points[y][1] &&
				points[x][0] < points[y][0])
	})

	ans, n := 0, len(points)
	for i := 0; i < n; {
		ans++
		var j int
		for j = i + 1; j < n && points[j][0] <= points[i][1]; j++ {}
		i = j
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
