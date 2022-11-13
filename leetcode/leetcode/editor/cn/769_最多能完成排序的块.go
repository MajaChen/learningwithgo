package leetcode

//leetcode submit region begin(Prohibit modification and deletion)
func maxChunksToSorted(arr []int) int {
	ans, max := 0, 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
		if max == i {
			ans++
		}
	}
	return ans
}
//leetcode submit region end(Prohibit modification and deletion)
