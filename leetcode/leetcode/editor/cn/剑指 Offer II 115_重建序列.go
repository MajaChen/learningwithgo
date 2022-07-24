package leetcode

//leetcode submit region begin(Prohibit modification and deletion)
func sequenceReconstruction(nums []int, sequences [][]int) bool {
	in, relations := make([]int, len(nums)+1), make([][]int, len(nums)+1)
	for _, sequence := range sequences {
		for i := 0; i < len(sequence)-1; i++ {
			x, y := sequence[i], sequence[i+1]
			in[y] += 1
			relations[x] = append(relations[x], y)
		}
	}

	q := make([]int, 0, len(nums))
	for i := 1; i < len(in); i++ {
		if in[i] == 0 {
			q = append(q, i)
		}
	}

	for len(q) > 0 {
		if len(q) > 1 {
			return false
		}
		i := q[0]
		q = q[:0]
		for _, j := range relations[i] {
			if in[j] -= 1; in[j] == 0 {
				q = append(q, j)
			}
		}
	}

	return true
}

//leetcode submit region end(Prohibit modification and deletion)
