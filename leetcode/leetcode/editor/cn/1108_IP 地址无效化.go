package leetcode

import "regexp"

//leetcode submit region begin(Prohibit modification and deletion)
func defangIPaddr(address string) string {
	p, _ := regexp.Compile("\\.")
	return p.ReplaceAllLiteralString(address, "[.]")
}

//leetcode submit region end(Prohibit modification and deletion)
