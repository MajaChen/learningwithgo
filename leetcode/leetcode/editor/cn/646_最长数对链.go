package leetcode

import (
	"math"
	"sort"
)

//leetcode submit region begin(Prohibit modification and deletion)
func findLongestChain(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] < pairs[j][1]
	})

	ans, s := 0, math.MinInt32
	for _, p := range pairs {
		if p[0] > s {
			ans++
			s = p[1]
		}
	}

	return ans
}
//leetcode submit region end(Prohibit modification and deletion)
