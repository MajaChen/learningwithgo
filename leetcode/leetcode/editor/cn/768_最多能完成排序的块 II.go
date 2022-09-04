package leetcode

//leetcode submit region begin(Prohibit modification and deletion)
func maxChunksToSorted(arr []int) int {
	s := make([]int, 0)
	for i := 0; i < len(arr); i++ {
		if len(s) == 0 || arr[i] >= s[len(s)-1] {
			s = append(s, arr[i])
		} else {
			m := s[len(s)-1]
			s = s[:len(s)-1]
			for ; len(s) > 0 && s[len(s)-1] > arr[i]; {
				s = s[:len(s)-1]
			}
			s = append(s, m)
		}
	}
	return len(s)
}
//leetcode submit region end(Prohibit modification and deletion)
