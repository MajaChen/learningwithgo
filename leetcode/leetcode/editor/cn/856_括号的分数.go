package leetcode

import . "learning/common"

//leetcode submit region begin(Prohibit modification and deletion)
func scoreOfParentheses(s string) int {
	stack := Stack{Elems: make([]interface{}, 0)}
	stack.Push(0)
	for _, r := range s {
		if r == 40 {
			stack.Push(0)
		} else if r == 41 {
			i := stack.Pop().(int)
			if i *= 2; i < 1 {
				i = 1
			}
			j := stack.Pop().(int)
			j += i
			stack.Push(j)
		}
	}
	return stack.Pop().(int)
}

//leetcode submit region end(Prohibit modification and deletion)
