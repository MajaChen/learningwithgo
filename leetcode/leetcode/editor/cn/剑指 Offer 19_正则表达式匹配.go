package leetcode

//leetcode submit region begin(Prohibit modification and deletion)

func isMatch(s string, p string) bool {
	ls, lp := len(s), len(p)
	matchMapping := make([][]bool, 0, ls+1)
	for i := 0; i < ls+1; i++ {
		matchMapping = append(matchMapping, make([]bool, lp+1))
	}

	// matchMapping[i][j] indicates if s[i-1] match s[j-1]
	for i := 0; i <= ls; i++ {
		for j := 0; j <= lp; j++ {
			// empty pattern
			if j == 0 {
				matchMapping[i][j] = i == 0
			} else { // non-empty pattern
				// non-with * spell
				if p[j-1:j] != "*" {
					if i > 0 && (s[i-1:i] == p[j-1:j] || p[j-1:j] == ".") {
						matchMapping[i][j] = matchMapping[i-1][j-1]
					}
				} else { // with * spell
					// * doesn't match
					if j >= 2 {
						matchMapping[i][j] = matchMapping[i][j-2]
					}

					// * does match
					if i >= 1 && j >= 2 && (s[i-1:i] == p[j-2:j-1] || p[j-2:j-1] == ".") {
						matchMapping[i][j] = matchMapping[i-1][j]
					}
				}
			}
		}
	}
	return matchMapping[ls][lp]
}

//leetcode submit region end(Prohibit modification and deletion)
