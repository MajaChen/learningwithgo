package leetcode

import (
	"math"
	"sort"
)

//leetcode submit region begin(Prohibit modification and deletion)

var iniPos = -int(math.Pow(10, 5))

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(x, y int) bool {
		return intervals[x][1] <= intervals[y][1]
	})
	pos, ans := iniPos, 0
	for _, interval := range intervals {
		if interval[0] >= pos {
			ans++
			pos = interval[1]
		}
	}
	return len(intervals) - ans
}

//leetcode submit region end(Prohibit modification and deletion)
