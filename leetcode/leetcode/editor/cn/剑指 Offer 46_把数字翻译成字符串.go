package leetcode

import (
	"strconv"
	"strings"
)

//leetcode submit region begin(Prohibit modification and deletion)
func translateNum(num int) int {
	s := strconv.Itoa(num)
	n, ans := len(s), 1
	m := make([]int, n+1)

	m[n-1], m[n] = 1, 1
	for i := n - 2; i >= 0; i-- {
		ans = 0
		// option a
		ans += m[i+1]
		// option b
		sub := s[i : i+2]
		j, _ := strconv.Atoi(sub)
		if j <= 25 && !strings.HasPrefix(sub, "0") {
			ans += m[i+2]
		}
		m[i] = ans
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
