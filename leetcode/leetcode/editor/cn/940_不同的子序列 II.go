package leetcode

//leetcode submit region begin(Prohibit modification and deletion)
func distinctSubseqII(s string) int {
	ans := 0
	lastK := make([]int, 26)
	var m int = 1e9 + 7
	for i := 0; i < 26; i++ {
		lastK[i] = -1
	}
	f := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		f[i] = 1
	}
	for i, c := range s {
		for _, j := range lastK {
			if j != -1 {
				f[i] = (f[i]+f[j]) % m
			}
		}
		lastK[c -'a'] = i
	}
	for _, j := range lastK {
		if j != -1 {
			ans = (ans + f[j]) % m
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
