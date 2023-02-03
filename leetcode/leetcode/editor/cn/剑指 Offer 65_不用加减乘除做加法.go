package leetcode

//leetcode submit region begin(Prohibit modification and deletion)
func add(a int, b int) int {
	for b != 0 {
		c := (a & b) << 1
		a = a ^ b
		b = c
	}
	return a
}

//leetcode submit region end(Prohibit modification and deletion)
