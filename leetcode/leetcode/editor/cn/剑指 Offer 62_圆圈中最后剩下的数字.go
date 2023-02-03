package leetcode

//leetcode submit region begin(Prohibit modification and deletion)

func lastRemainingRe(n, m int) int {
	if n == 1 {
		return 0
	}
	x := lastRemainingRe(n-1, m)
	return (m + x) % n
}

func lastRemaining(n int, m int) int {
	return lastRemainingRe(n, m)
}

//leetcode submit region end(Prohibit modification and deletion)
