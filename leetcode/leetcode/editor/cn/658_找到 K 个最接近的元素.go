package leetcode

import (
	"math"
	"sort"
)

//leetcode submit region begin(Prohibit modification and deletion)
func findClosestElements(arr []int, k int, x int) []int {
	sort.Slice(arr, func(i, j int) bool {
		if math.Abs(float64(arr[i]-x)) == math.Abs(float64(arr[j]-x)) {
			return arr[i] < arr[j]
		} else {
			return math.Abs(float64(arr[i]-x)) < math.Abs(float64(arr[j]-x))
		}
	})
	ans := arr[:k]
	sort.Ints(ans)
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
