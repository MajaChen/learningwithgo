package leetcode

//leetcode submit region begin(Prohibit modification and deletion)
func maxProfit(prices []int) int {
	maximalProfit, l := 0, 0
	for i := 1; i < len(prices); i++ {
		if prices[i] >= prices[l] {
			if prices[i]-prices[l] > maximalProfit {
				maximalProfit = prices[i] - prices[l]
			}
		} else {
			l = i
		}
	}
	return maximalProfit
}

//leetcode submit region end(Prohibit modification and deletion)
