package leetcode

import (
	"sort"
)

//leetcode submit region begin(Prohibit modification and deletion)

func f(x int) int {
	c := 0
	i := 5
	for x >= i {
		c += x / i
		i *= 5
	}
	return c
}

// 在整数空间范围内，求阶乘末尾0的个数>= k的最小整数
func g(k int) int {
	return sort.Search(5*k, func(i int) bool { return f(i) >= k })
}

func preimageSizeFZF(k int) int {
	return g(k+1) - g(k)
}

//leetcode submit region end(Prohibit modification and deletion)
