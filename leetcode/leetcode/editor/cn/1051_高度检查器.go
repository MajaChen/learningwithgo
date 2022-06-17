package leetcode

//leetcode submit region begin(Prohibit modification and deletion)
func heightChecker(heights []int) int {
	ans, counters := 0, make([]int, 101)
	for _, height := range heights {
		counters[height]++
	}

	for i, j := 1, 0; i < len(counters); i++ {
		for ; counters[i] > 0; counters[i]-- {
			if i != heights[j] {
				ans++
			}
			j++
		}
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
