package leetcode

//leetcode submit region begin(Prohibit modification and deletion)

func spiralOrder(matrix [][]int) []int {

	if len(matrix) == 0 {
		return []int{}
	}

	i, j, k, ans := 0, -1, 0, make([]int, 0, len(matrix)*len(matrix[0]))
	l, r, u, d := 0, len(matrix[0])-1, 0, len(matrix)-1
	functionsMapping := map[int]func(){
		0: func() { // to right
			if j = j + 1; j >= r {
				k, u = (k+1)%4, u+1
			}
		},
		1: func() { // down
			if i = i + 1; i >= d {
				k, r = (k+1)%4, r-1
			}
		},
		2: func() { // to left
			if j = j - 1; j <= l {
				k, d = (k+1)%4, d-1
			}
		},
		3: func() { // up
			if i = i - 1; i <= u {
				k, l = (k+1)%4, l+1
			}
		}}

	for l <= r && u <= d {
		functionsMapping[k]()
		ans = append(ans, matrix[i][j])
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
