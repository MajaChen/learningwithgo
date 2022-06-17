package leetcode

//leetcode submit region begin(Prohibit modification and deletion)
func findDiagonalOrder(mat [][]int) []int {
	i, j, k, m, n := 1, -1, 0, len(mat), len(mat[0])
	ans := make([]int, 0, m*n)
	functionsMapping := map[int]func(){
		0: func() { // up
			if i == 0 || j == n-1 {
				if j == n-1 {
					i++
				} else if i == 0 {
					j++
				}
				k = 1
			} else {
				i--
				j++
			}
		},
		1: func() { // down
			if j == 0 || i == m-1 {
				if i == m-1 {
					j++
				} else if j == 0 {
					i++
				}
				k = 0
			} else {
				i++
				j--
			}
		}}

	for i != m-1 || j != n-1 {
		functionsMapping[k]()
		ans = append(ans, mat[i][j])
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
