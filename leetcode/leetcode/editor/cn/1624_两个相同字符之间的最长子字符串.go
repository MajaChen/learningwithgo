package leetcode

//leetcode submit region begin(Prohibit modification and deletion)
func maxLengthBetweenEqualCharacters(s string) int {
	mapping, ans := make(map[uint8]int), -1
	for i := 0; i < len(s); i++ {
		if _, ok := mapping[s[i]]; ok {
			if l := i - mapping[s[i]] - 1; l > ans {
				ans = l
			}
		} else {
			mapping[s[i]] = i
		}
	}

	return ans

}

//leetcode submit region end(Prohibit modification and deletion)
