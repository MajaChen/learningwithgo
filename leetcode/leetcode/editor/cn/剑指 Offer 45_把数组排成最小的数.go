package leetcode

import (
	"sort"
	"strconv"
	"strings"
)

//leetcode submit region begin(Prohibit modification and deletion)
func isSmaller(sa, sb string) bool {
	for i := 0; i < len(sa); i++ {
		if sa[i] < sb[i] {
			return true
		} else if sa[i] > sb[i] {
			return false
		}
	}
	return false
}

type sortableStrs []string

func (s sortableStrs) Len() int {
	return len(s)
}

func (s sortableStrs) Less(i, j int) bool {
	if isSmaller(s[i]+s[j], s[j]+s[i]) {
		return true
	} else {
		return false
	}
}

func (s sortableStrs) Swap(i, j int) {
	t := s[i]
	s[i] = s[j]
	s[j] = t
}

func minNumber(nums []int) string {
	ss := make([]string, 0, len(nums))
	for _, num := range nums {
		ss = append(ss, strconv.Itoa(num))
	}

	sort.Sort(sortableStrs(ss))
	var builder strings.Builder
	for _, str := range ss {
		builder.WriteString(str)
	}
	return builder.String()
}

//leetcode submit region end(Prohibit modification and deletion)
