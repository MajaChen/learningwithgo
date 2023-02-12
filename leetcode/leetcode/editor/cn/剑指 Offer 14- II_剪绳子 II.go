package leetcode

//leetcode submit region begin(Prohibit modification and deletion)
func cuttingRope(n int) int {
	if n <= 3 {
		return n - 1
	}
	var ret int64 = 1
	if n%3 == 1 {
		ret = 4
		n = n - 4
	}
	if n%3 == 2 {
		ret = 2
		n = n - 2
	}
	for n > 0 {
		ret = ret * 3 % 1000000007
		n -= 3
	}
	return int(ret)

}

//leetcode submit region end(Prohibit modification and deletion)
