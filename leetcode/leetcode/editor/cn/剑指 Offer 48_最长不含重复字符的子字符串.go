package leetcode

//leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	mapping, l, r, maxLen := make(map[byte]int), 0, 0, 1
	mapping[s[0]] = 0
	for i := 1; i < len(s); i++ {
		r = i
		if _, ok := mapping[s[i]]; ok {
			var index int
			for index = l; index < mapping[s[i]]; index++ {
				delete(mapping, s[index])
			}
			l = mapping[s[i]] + 1
		}
		mapping[s[i]] = i

		if r-l+1 > maxLen {
			maxLen = r - l + 1
		}
	}
	return maxLen
}

//leetcode submit region end(Prohibit modification and deletion)
