package leetcode

//leetcode submit region begin(Prohibit modification and deletion)
var counter int

func totalNQueensRe(n int, row, col, l, r int) {

	if row >= n {
		counter++
		return
	}

	bits := (^(col | l | r)) & ((1 << n) - 1)
	for bits > 0 {
		p := bits & (-bits)
		bits = bits & (bits - 1)
		totalNQueensRe(n, row+1, col|p, (l|p)<<1, (r|p)>>1)
	}
	return
}

func totalNQueens(n int) int {

	counter = 0
	totalNQueensRe(n, 0, 0, 0, 0)
	return counter
}

//leetcode submit region end(Prohibit modification and deletion)
