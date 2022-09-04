package leetcode

import "strings"

//leetcode submit region begin(Prohibit modification and deletion)
func stringMatching(words []string) []string {
	ans := make([]string, 0, len(words))
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words); j++ {
			if i == j {
				continue
			}
			if strings.Index(words[j], words[i]) >= 0 {
				ans = append(ans, words[i])
				break
			}
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
