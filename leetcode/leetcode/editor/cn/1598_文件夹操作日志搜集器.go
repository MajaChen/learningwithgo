package leetcode

import . "learning/common"

//leetcode submit region begin(Prohibit modification and deletion)

func minOperations(logs []string) int {
	s := Stack{Elems: make([]interface{}, 0)}
	for _, log := range logs {
		if log == "./" {
		} else if log == "../" {
			if !s.IsEmpty() {
				s.Pop()
			}
		} else {
			s.Push(0)
		}
	}
	return s.Len()
}

//leetcode submit region end(Prohibit modification and deletion)
