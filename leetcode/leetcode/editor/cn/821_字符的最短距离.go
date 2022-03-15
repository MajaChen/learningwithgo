package leetcode

import "math"

//leetcode submit region begin(Prohibit modification and deletion)

var (
	magicNum = int(math.Pow(10, 4))
)

func shortestToChar(s string, c byte) []int {

	ans := make([]int, len(s))
	for i, l, r := 0, magicNum, magicNum; i < len(s); i++ {
		if s[i] == c {
			ans[i] = 0
			r = magicNum
			l = i
		} else {
			if l == magicNum {
				var j int
				for j = i - 1; j >= 0 && s[j] != c; j-- {
				}
				if j >= 0 {
					l = j
				}
			}
			if r == magicNum {
				var j int
				for j = i + 1; j < len(s) && s[j] != c; j++ {
				}
				if j < len(s) {
					r = j
				}
			}
			ans[i] = int(math.Min(math.Abs(float64(i-l)), math.Abs(float64(r-i))))
		}
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
