package leetcode

import "math"

//leetcode submit region begin(Prohibit modification and deletion)
func maxProduct(nums []int) int {
	first, second := math.MinInt32, math.MinInt32
	for _, num := range nums {
		if num > first {
			second = first
			first = num
			continue
		}
		if num > second {
			second = num
		}
	}
	return (first-1)*(second-1)
}
//leetcode submit region end(Prohibit modification and deletion)
