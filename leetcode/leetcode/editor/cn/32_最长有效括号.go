package leetcode

import . "learning/common"

//leetcode submit region begin(Prohibit modification and deletion)
func longestValidParentheses(s string) int {

	stack, longestLength := Stack{Elems: make([]interface{}, 0)}, 0
	stack.Push(-1)
	for i := 0; i < len(s); i++ {
		if s[i] == 40 {
			stack.Push(i)
		} else {
			stack.Pop()
			if stack.IsEmpty() {
				stack.Push(i)
				continue
			}
			if j := stack.GetTop().(int); i-j > longestLength {
				longestLength = i - j
			}
		}
	}

	return longestLength
}

//leetcode submit region end(Prohibit modification and deletion)
