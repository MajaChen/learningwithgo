package leetcode

//leetcode submit region begin(Prohibit modification and deletion)

func deleteFromMapping(s string, mapping map[string]int) {

	mapping[s] = mapping[s] - 1
	if count, _ := mapping[s]; count <= 0 {
		delete(mapping, s)
	}
}

func addToMapping(s string, mapping map[string]int) {

	if _, ok := mapping[s]; ok {
		mapping[s] = mapping[s] + 1
	} else {
		mapping[s] = 1
	}
}

func maxFreq(s string, maxLetters int, minSize int, maxSize int) int {

	mapping, characterSet := make(map[string]int), make(map[string]int)
	for i := 0; i < minSize-1; i++ {
		addToMapping(s[i:i+1], characterSet)
	}
	for i := 0; i+minSize <= len(s); i++ {
		leftBound, rightBound := i+minSize-1, i+maxSize-1
		var j int
		for j = leftBound; j <= rightBound && j < len(s); j++ {
			addToMapping(s[j:j+1], characterSet)
			if len(characterSet) <= maxLetters {
				if _, ok := mapping[s[i:j+1]]; ok {
					mapping[s[i:j+1]] = mapping[s[i:j+1]] + 1
				} else {
					mapping[s[i:j+1]] = 1
				}
			}
		}

		deleteFromMapping(s[i:i+1], characterSet)
		for j -= 1; j > leftBound; j-- {
			deleteFromMapping(s[j:j+1], characterSet)
		}
	}

	var maxC int
	for _, c := range mapping {
		if c > maxC {
			maxC = c
		}
	}

	return maxC
}

//leetcode submit region end(Prohibit modification and deletion)
