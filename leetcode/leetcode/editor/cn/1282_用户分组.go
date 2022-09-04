package leetcode

//leetcode submit region begin(Prohibit modification and deletion)

func groupThePeople(groupSizes []int) [][]int {
	n := len(groupSizes)
	matrix := make([][]int, n+1)
	for i := 0; i < n; i++ {
		matrix = append(matrix, make([]int, 0))
	}
	ans := make([][]int, 0)
	for i, j := 0, 0; i < n; i++ {
		j = groupSizes[i]
		matrix[j] = append(matrix[j], i)
		if len(matrix[j]) == j {
			ans = append(ans, matrix[j])
			matrix[j] = make([]int, 0)
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
