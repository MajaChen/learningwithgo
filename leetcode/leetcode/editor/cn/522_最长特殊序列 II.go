package leetcode

//leetcode submit region begin(Prohibit modification and deletion)

func isSubStr(s, t string) bool {
	var i, j int
	for i, j = 0, 0; i < len(s) && j < len(t); {
		if s[i] == t[j] {
			i++
			j++
		} else {
			j++
		}
	}

	return i == len(s)
}

func findLUSlength(strs []string) int {
	n, ans := len(strs), -1
next:
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i != j && isSubStr(strs[i], strs[j]) {
				continue next
			}
		}
		if len(strs[i]) > ans {
			ans = len(strs[i])
		}
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
