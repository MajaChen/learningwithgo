package leetcode

import "sort"

//leetcode submit region begin(Prohibit modification and deletion)
func findContentChildren(g []int, s []int) int {

	sort.Ints(g)
	sort.Ints(s)

	var counter int
	for i, j := 0, 0; i < len(s); i++ {
		if j < len(g) && g[j] <= s[i] {
			j++
			counter++
		}
	}

	return counter
}

//leetcode submit region end(Prohibit modification and deletion)
