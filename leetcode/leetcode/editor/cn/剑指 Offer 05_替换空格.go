package leetcode

import "strings"

//leetcode submit region begin(Prohibit modification and deletion)
func replaceSpace(s string) string {

	builder := &strings.Builder{}
	builder.Grow(int(float32(len(s)) * 1.2))
	for i := 0; i < len(s); i++ {
		if s[i:i+1] == " " {
			builder.WriteString("%20")
		} else {
			builder.WriteString(s[i : i+1])
		}
	}
	return builder.String()
}

//leetcode submit region end(Prohibit modification and deletion)
