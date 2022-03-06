package leetcode

import (
	"math"
)

//leetcode submit region begin(Prohibit modification and deletion)
func maximumSum(arr []int) int {

	// 分别表示：漏掉一个取得的最大值，一个不漏取得的最大值
	amax, bmax, maxSum := arr[0], arr[0], arr[0]
	for i := 1; i < len(arr); i++ {
		// 不漏掉 arr[i]：之前漏掉一个取得的最大值 vs 漏掉arr[i]：之前一个都不漏掉取得的最大值
		if amax = int(math.Max(float64(amax+arr[i]), float64(bmax))); amax > maxSum {
			maxSum = amax
		}
		if bmax = int(math.Max(float64(bmax+arr[i]), float64(arr[i]))); bmax > maxSum {
			maxSum = bmax
		}
	}

	return maxSum
}

//leetcode submit region end(Prohibit modification and deletion)
