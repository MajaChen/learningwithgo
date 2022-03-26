package leetcode

//leetcode submit region begin(Prohibit modification and deletion)
func findRepeatNumber(nums []int) int {

	arr := make([]int, len(nums))
	for _, num := range nums {
		if arr[num]++; arr[num] > 1 {
			return num
		}
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)
