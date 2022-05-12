package leetcode

import "math"

//leetcode submit region begin(Prohibit modification and deletion)
func maxSubArray(nums []int) int {
	sum, minSum, maximalSubArray := make([]int, len(nums)+1), 0, math.MinInt32
	for i := 1; i <= len(nums); i++ {
		sum[i] = sum[i-1] + nums[i-1]
		if sum[i]-minSum > maximalSubArray {
			maximalSubArray = sum[i] - minSum
		}
		if sum[i] < minSum {
			minSum = sum[i]
		}
	}
	return maximalSubArray
}

//leetcode submit region end(Prohibit modification and deletion)
