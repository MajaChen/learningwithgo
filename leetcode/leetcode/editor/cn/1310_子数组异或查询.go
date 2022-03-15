package leetcode

//leetcode submit region begin(Prohibit modification and deletion)
func xorQueries(arr []int, queries [][]int) []int {

	xor := make([]int, 0, len(arr))
	xor = append(xor, arr[0])
	for i := 1; i < len(arr); i++ {
		xor = append(xor, xor[i-1]^arr[i])
	}

	res := make([]int, 0, len(queries))
	for i := 0; i < len(queries); i++ {
		var j int
		if queries[i][0] == 0 {
			j = xor[queries[i][1]]
		} else {
			j = xor[queries[i][0]-1] ^ xor[queries[i][1]]
		}
		res = append(res, j)
	}

	return res
}

//leetcode submit region end(Prohibit modification and deletion)
