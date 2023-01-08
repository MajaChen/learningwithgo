package leetcode

//leetcode submit region begin(Prohibit modification and deletion)
func partitionLabels(s string) []int {
	m := make(map[int32]int)
	for i, r := range s {
		m[r] = i
	}

	ans := make([]int, 0)
	for start, end := 0, m[int32(s[0])]; start < len(s); {
		for i := start; i <= end; i++ {
			if m[int32(s[i])] > end {
				end = m[int32(s[i])]
			}
		}
		ans = append(ans, end-start+1)
		if start = end + 1; start < len(s) {
			end = m[int32(s[start])]
		}

	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
