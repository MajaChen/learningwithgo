package leetcode

//leetcode submit region begin(Prohibit modification and deletion)
func minCostToMoveChips(position []int) int {
	i, j := 0, 0
	for _, p := range position {
		if p % 2 == 0 {
			i++
		} else {
			j++
		}
	}
	if i > j {
		return j
	} else {
		return i
	}
}
//leetcode submit region end(Prohibit modification and deletion)
