package leetcode

import (
	"sort"
)

//leetcode submit region begin(Prohibit modification and deletion)
func trimMean(arr []int) float64 {
	sort.Ints(arr)
	totalSum, n := 0, len(arr)
	start, end := n/20, n-n/20
	for i := start; i < end; i++ {
		totalSum += arr[i]
	}
	return float64(totalSum)/float64(n-n/20*2)
}
//leetcode submit region end(Prohibit modification and deletion)
