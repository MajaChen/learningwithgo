package leetcode

//leetcode submit region begin(Prohibit modification and deletion)

func numEnclavesTraversal(grid [][]int, i, j, m, n int) {

	if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] != 1 {
		return
	}

	grid[i][j] = 2
	numEnclavesTraversal(grid, i-1, j, m, n)
	numEnclavesTraversal(grid, i+1, j, m, n)
	numEnclavesTraversal(grid, i, j-1, m, n)
	numEnclavesTraversal(grid, i, j+1, m, n)
}

func numEnclaves(grid [][]int) int {

	m, n := len(grid), len(grid[0])
	for j := 0; j < n; j++ {
		numEnclavesTraversal(grid, 0, j, m, n)
		numEnclavesTraversal(grid, m-1, j, m, n)
	}

	for i := 0; i < m; i++ {
		numEnclavesTraversal(grid, i, 0, m, n)
		numEnclavesTraversal(grid, i, n-1, m, n)
	}

	var count int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				count++
			} else if grid[i][j] == 2 {
				grid[i][j] = 1
			}
		}
	}

	return count
}

//leetcode submit region end(Prohibit modification and deletion)
