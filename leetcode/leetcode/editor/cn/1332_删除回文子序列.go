package leetcode

//leetcode submit region begin(Prohibit modification and deletion)
func removePalindromeSub(s string) int {

	if len(s) == 0 {
		return 0
	}

	for i, j := 0, len(s)-1; i < j; {
		if s[i] != s[j] {
			return 2
		}
		i++
		j--
	}

	return 1

}

//leetcode submit region end(Prohibit modification and deletion)
