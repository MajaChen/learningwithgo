package leetcode

import "regexp"

//leetcode submit region begin(Prohibit modification and deletion)

func str2Int(str string) int {

	return 0
}

func strToInt(str string) int {

	var index []int
	if index = regexp.MustCompile("\\S").FindStringIndex(str); index == nil {
		return 0
	}
	var findString string
	if findString = regexp.MustCompile("[\\+\\-]?[0-9]+").FindString(str[index[0]:]); len(findString) == 0 {
		return 0
	}

	return str2Int(findString)
}

//leetcode submit region end(Prohibit modification and deletion)
