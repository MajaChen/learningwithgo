package leetcode

import "strconv"

//leetcode submit region begin(Prohibit modification and deletion)
func maximumSwap(num int) int {
	s := []rune(strconv.Itoa(num))
	for i := 0; i < len(s); i++ {
		m := -1
		for j := i + 1; j < len(s); j++ {
			if s[j] > s[i] {
				if m < 0 {
					m = j
					continue
				}
				if s[j] >= s[m] {
					m = j
				}
			}
		}
		if m >= 0 {
			t := s[i]
			s[i] = s[m]
			s[m] = t
			num, _ = strconv.Atoi(string(s))
			return num
		}
	}
	return num
}

//leetcode submit region end(Prohibit modification and deletion)
