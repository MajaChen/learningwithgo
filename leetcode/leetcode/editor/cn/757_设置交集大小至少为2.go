package leetcode

import "sort"

//leetcode submit region begin(Prohibit modification and deletion)
func intersectionSizeTwo(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0] ||
			(intervals[i][0] == intervals[j][0] &&
				intervals[i][1] > intervals[j][1])
	})

	ans, m, n, counters := 0, 2, len(intervals), make([]int, len(intervals))
	for i := n - 1; i >= 0; i-- {
		for j, k := intervals[i][0], counters[i]; k < m; k++ {
			ans++
			for p := i - 1; p >= 0 && intervals[p][1] >= j; p-- {
				counters[p] += 1
			}
			j++
		}
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
