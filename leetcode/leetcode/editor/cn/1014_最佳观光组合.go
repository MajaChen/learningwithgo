package leetcode

//leetcode submit region begin(Prohibit modification and deletion)
func maxScoreSightseeingPair(values []int) int {

	var maximalScore int
	pre := values[0] + 0
	for i := 1; i < len(values); i++ {
		if pre+values[i]-i > maximalScore {
			maximalScore = pre + values[i] - i
		}
		if values[i]+i > pre {
			pre = values[i] + i
		}
	}
	return maximalScore
}

//leetcode submit region end(Prohibit modification and deletion)
