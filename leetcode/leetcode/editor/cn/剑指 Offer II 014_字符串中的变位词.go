package leetcode

//leetcode submit region begin(Prohibit modification and deletion)
func checkInclusion(s1 string, s2 string) bool {
	l1, l2 := len(s1), len(s2)
	if l1 > l2 {
		return false
	}

	var a1, a2 [26]int
	for i := 0; i < l1; i++ {
		a1[s1[i]-'a']++
		a2[s2[i]-'a']++
	}
	if a1 == a2 {
		return true
	}

	for i := 0; i < l2-l1; i++ {
		a2[s2[i]-'a']--
		a2[s2[i+l1]-'a']++
		if a1 == a2 {
			return true
		}
	}

	return false
}

//leetcode submit region end(Prohibit modification and deletion)
