package leetcode

//leetcode submit region begin(Prohibit modification and deletion)
func firstUniqChar(s string) byte {
	chars := [26]int{}
	for _, r := range s {
		chars[r-'a']++
	}
	for i, r := range s {
		if chars[r-'a'] == 1 {
			return s[i]
		}
	}
	return ' '
}

//leetcode submit region end(Prohibit modification and deletion)
