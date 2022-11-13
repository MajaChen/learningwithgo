package leetcode

//leetcode submit region begin(Prohibit modification and deletion)

func traverseIslands(grid [][]int, i, j int) int {

	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || grid[i][j] != 1 {
		return 0
	}

	grid[i][j] = 2
	counter := 1
	counter += traverseIslands(grid, i+1, j)
	counter += traverseIslands(grid, i-1, j)
	counter += traverseIslands(grid, i, j+1)
	counter += traverseIslands(grid, i, j-1)
	return counter
}

func largestIsland(grid [][]int) int {
	ans := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			if num := traverseIslands(grid, i, j); num > ans {
				ans = num
			}
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
