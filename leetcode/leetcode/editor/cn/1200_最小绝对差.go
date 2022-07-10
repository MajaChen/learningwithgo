package leetcode

import (
	"math"
	"sort"
)

//leetcode submit region begin(Prohibit modification and deletion)
func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	minimalDiff, ans := math.MaxInt32, make([][]int, 0)
	for i := 1; i < len(arr); i++ {
		if diff := arr[i] - arr[i-1]; diff < minimalDiff {
			minimalDiff = diff
			ans = ans[:0]
			ans = append(ans, []int{arr[i-1], arr[i]})
		} else if diff == minimalDiff {
			ans = append(ans, []int{arr[i-1], arr[i]})
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
