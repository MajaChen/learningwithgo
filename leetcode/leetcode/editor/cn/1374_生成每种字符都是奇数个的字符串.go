package leetcode

import "strings"


//leetcode submit region begin(Prohibit modification and deletion)

var magicCharacter = "a"
var magicCharacterPro = "b"

func generateTheStringInternal(n int, c string) string {
	sb := strings.Builder{}
	for i := 0; i < n; i++ {
		sb.WriteString(c)
	}
	return sb.String()
}

func generateTheString(n int) string {

	if n % 2 == 0 {
		return generateTheStringInternal(n-1, magicCharacter) +
			generateTheStringInternal(1, magicCharacterPro)
	} else {
		return generateTheStringInternal(n, magicCharacter)
	}
}
//leetcode submit region end(Prohibit modification and deletion)
