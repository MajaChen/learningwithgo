package leetcode

//leetcode submit region begin(Prohibit modification and deletion)

const (
	whatever = 0
)

func findPairs(nums []int, k int) int {
	mappingA, mappingB := make(map[int]int), make(map[int]int)
	for _, num := range nums {
		if _, ok := mappingA[num-k]; ok {
			if _, ok := mappingB[num-k]; !ok {
				mappingB[num-k] = whatever
			}
		}
		if _, ok := mappingA[num+k]; ok {
			if _, ok := mappingB[num]; !ok {
				mappingB[num] = whatever
			}
		}

		mappingA[num] = 0
	}
	return len(mappingB)
}

//leetcode submit region end(Prohibit modification and deletion)
