package leetcode

//leetcode submit region begin(Prohibit modification and deletion)

func matchPattern(word, pattern string) bool {
	runesA, runesB := []rune(word), []rune(pattern)
	mappingA, mappingB := make(map[rune]rune), make(map[rune]rune)
	for i := 0; i < len(word); i++ {
		ra, rb := runesA[i], runesB[i]
		if r, ok := mappingA[ra]; ok {
			if r != rb {
				return false
			}
		} else {
			mappingA[ra] = rb
		}
		if r, ok := mappingB[rb]; ok {
			if r != ra {
				return false
			}
		} else {
			mappingB[rb] = ra
		}
	}
	return true
}

func findAndReplacePattern(words []string, pattern string) []string {
	ans := make([]string, 0, len(words))
	for _, word := range words {
		if matchPattern(word, pattern) {
			ans = append(ans, word)
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
