package leetcode

import "math"

//leetcode submit region begin(Prohibit modification and deletion)
func minStartValue(nums []int) int {
	sum, minimal := 0, math.MaxInt32
	for i := 0; i < len(nums); i++ {
		if sum += nums[i]; sum < minimal {
			minimal = sum
		}
	}
	if minimal > 0 {
		return 1
	}
	return 1 - minimal
}

//leetcode submit region end(Prohibit modification and deletion)
