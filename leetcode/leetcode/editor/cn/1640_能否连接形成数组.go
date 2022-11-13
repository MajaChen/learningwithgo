package leetcode

//leetcode submit region begin(Prohibit modification and deletion)
func canFormArray(arr []int, pieces [][]int) bool {
	mapping := make(map[int]int)
	for index, piece := range pieces {
		mapping[piece[0]] = index
	}

	for i := 0; i < len(arr); {
		if _, ok := mapping[arr[i]]; !ok {
			return false
		}
		nums := pieces[mapping[arr[i]]]
		for _, num := range nums {
			if arr[i] != num {
				return false
			}
			i++
		}
	}

	return true
}
//leetcode submit region end(Prohibit modification and deletion)
