package leetcode

//leetcode submit region begin(Prohibit modification and deletion)
func sumNums(n int) int {
	ans := 0
	var sum func(int) bool
	sum = func(n int) bool {
		ans += n
		return n > 0 && sum(n-1)
	}
	sum(n)
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
