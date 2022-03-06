package leetcode

//leetcode submit region begin(Prohibit modification and deletion)

func lowCase(s uint8) bool {
	return s >= 97 && s <= 122
}

func canMatch(query, pattern string) bool {

	var i, j int
	for i, j = 0, 0; i < len(query) && j < len(pattern); {
		if query[i] == pattern[j] {
			i++
			j++
		} else if lowCase(query[i]) {
			i++
		} else {
			return false
		}
	}

	if j < len(pattern) {
		return false
	}

	if i < len(query) {
		for k := i; k < len(query); k++ {
			if !lowCase(query[k]) {
				return false
			}
		}
	}

	return true
}
func camelMatch(queries []string, pattern string) []bool {

	res := make([]bool, 0, len(queries))
	for i := 0; i < len(queries); i++ {
		if canMatch(queries[i], pattern) {
			res = append(res, true)
		} else {
			res = append(res, false)
		}
	}

	return res

}

//leetcode submit region end(Prohibit modification and deletion)
