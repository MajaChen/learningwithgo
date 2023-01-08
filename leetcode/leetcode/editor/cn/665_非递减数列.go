package leetcode

//leetcode submit region begin(Prohibit modification and deletion)

func isSort(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		x, y := nums[i], nums[i+1]
		if x > y {
			return false
		}
	}
	return true
}

func checkPossibility(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		x, y := nums[i], nums[i+1]
		if x > y {
			nums[i] = nums[i+1]
			if isSort(nums) {
				return true
			}
			nums[i] = x
			nums[i+1] = nums[i]
			return isSort(nums)
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
