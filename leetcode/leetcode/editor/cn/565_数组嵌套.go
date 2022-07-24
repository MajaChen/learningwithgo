package leetcode

//leetcode submit region begin(Prohibit modification and deletion)
func arrayNesting(nums []int) int {
	ans := 0
	for i := range nums {
		j, c, k := i, 0, i
		for j >= 0&&nums[j] >= 0 {
			c++
			k = j
			j = nums[j]
			nums[k] = -1
		}
		if c > ans {
			ans = c
		}
	}
	return ans
}
//leetcode submit region end(Prohibit modification and deletion)
