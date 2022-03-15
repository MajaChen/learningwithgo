package leetcode

//leetcode submit region begin(Prohibit modification and deletion)
func getSum(a int, b int) int {

	if a == 0 {
		return b
	}

	if b == 0 {
		return a
	}

	for b != 0 {
		carry := (a & b) << 1
		a = a ^ b
		b = carry
	}
	return a
}

//leetcode submit region end(Prohibit modification and deletion)
