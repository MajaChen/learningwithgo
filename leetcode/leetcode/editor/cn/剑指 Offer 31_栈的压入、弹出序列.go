package leetcode

import . "learning/common"

//leetcode submit region begin(Prohibit modification and deletion)

func validateStackSequences(pushed []int, popped []int) bool {
	stack := Stack{Elems: make([]interface{}, 0, 32)}
	for i, k := 0, 0; i < len(popped); i++ {
		if stack.IsEmpty() || stack.GetTop().(int) != popped[i] {
			for ; k < len(pushed) && pushed[k] != popped[i]; k++ {
				stack.Push(pushed[k])
			}
			k++
		} else {
			stack.Pop()
		}
	}
	return stack.IsEmpty()
}

//leetcode submit region end(Prohibit modification and deletion)
