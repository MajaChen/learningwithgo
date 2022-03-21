package leetcode

import (
	"math"
	"strings"
)

//leetcode submit region begin(Prohibit modification and deletion)

var queens [][]string
var template []string

func template2String(index int) string {

	template[index] = "Q"
	s := strings.Join(template, "")
	template[index] = "."
	return s
}

func solveNQueensRe(n int, row, col, l, r int, queen []string) {

	if row >= n {
		arr := make([]string, len(queen))
		copy(arr, queen)
		queens = append(queens, arr)
		return
	}

	bits := (^(col | l | r)) & ((1 << n) - 1)
	for bits > 0 {
		p := bits & (-bits)
		index := n - (int(math.Log(float64(p))/math.Log(2.0)) + 1)
		queen = append(queen, template2String(index))
		bits = bits & (bits - 1)
		solveNQueensRe(n, row+1, col|p, (l|p)<<1, (r|p)>>1, queen)
		queen = queen[:len(queen)-1]
	}
}

func solveNQueens(n int) [][]string {

	queens = make([][]string, 0)
	template = make([]string, 0, n)
	for i := 0; i < n; i++ {
		template = append(template, ".")
	}

	solveNQueensRe(n, 0, 0, 0, 0, []string{})
	return queens
}

//leetcode submit region end(Prohibit modification and deletion)
