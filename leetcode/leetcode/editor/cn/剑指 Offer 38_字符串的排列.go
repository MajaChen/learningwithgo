package leetcode

//leetcode submit region begin(Prohibit modification and deletion)

var (
	uniqueSet map[string]interface{}
)

func permutationWithRecursion(runes, partialResult []rune, startPoint int, visitedMapping []bool) {
	if startPoint >= len(runes) {
		result := make([]rune, len(runes))
		copy(result, partialResult)
		uniqueSet[string(result)] = nil
		return
	}

	for index, isVisited := range visitedMapping {
		if !isVisited {
			visitedMapping[index] = true
			partialResult = append(partialResult, runes[index])
			permutationWithRecursion(runes, partialResult, startPoint+1, visitedMapping)
			visitedMapping[index] = false
			partialResult = partialResult[:len(partialResult)-1]
		}
	}
}
func permutation(s string) []string {
	uniqueSet = make(map[string]interface{})
	permutationWithRecursion([]rune(s), make([]rune, 0, len(s)), 0, make([]bool, len(s)))
	results := make([]string, 0, len(uniqueSet))
	for k := range uniqueSet {
		results = append(results, k)
	}
	return results
}

//leetcode submit region end(Prohibit modification and deletion)
