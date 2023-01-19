package leetcode

//leetcode submit region begin(Prohibit modification and deletion)
func twoSum(nums []int, target int) []int {
	for i, j := 0, len(nums)-1; i < j; {
		if ai, aj := nums[i], nums[j]; ai+aj == target {
			return []int{nums[i], nums[j]}
		} else if ai+aj > target {
			j--
		} else {
			i++
		}
	}
	return []int{}
}

//leetcode submit region end(Prohibit modification and deletion)
