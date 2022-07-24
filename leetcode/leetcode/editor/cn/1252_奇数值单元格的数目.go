package leetcode

//leetcode submit region begin(Prohibit modification and deletion)
func oddCells(m int, n int, indices [][]int) int {
	ma, na := make([]int, m), make([]int, n)
	for _, indice := range indices {
		i, j := indice[0], indice[1]
		ma[i]++
		na[j]++
	}

	o, e := 0, 0
	for i := 0; i < n; i++ {
		if na[i] % 2 == 0 {
			e++
		} else {
			o++
		}
	}

	ans := 0
	for j := 0; j < m; j++ {
		if ma[j] % 2 == 0 {
			ans += o
		} else {
			ans += e
		}
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
