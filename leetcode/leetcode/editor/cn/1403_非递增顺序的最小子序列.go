package leetcode

import "sort"

//leetcode submit region begin(Prohibit modification and deletion)
func minSubsequence(nums []int) []int {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	var i, c, s int
	for i = 0; i < len(nums); i++ {
		s += nums[i]
	}
	for i = 0; i < len(nums); i++ {
		if c += nums[i]; c > s-c {
			break
		}
	}
	return nums[:i+1]
}

//leetcode submit region end(Prohibit modification and deletion)
