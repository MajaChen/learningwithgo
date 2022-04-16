package leetcode

import . "learning/common"

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func nextLargerNodes(head *ListNode) []int {

	mapping, stack, nodeCounts, count := make(map[int]int), &Stack{Elems: make([]interface{}, 0)}, make(map[int]int), 1
	for p := head; p != nil; p = p.Next {
		mapping[count] = p.Val
		for !stack.IsEmpty() && mapping[stack.GetTop().(int)] < p.Val {
			nodeCounts[stack.Pop().(int)] = p.Val
		}
		stack.Push(count)
		count++
	}

	res := make([]int, count-1)
	for index, count := range nodeCounts {
		res[index-1] = count
	}

	return res
}

//leetcode submit region end(Prohibit modification and deletion)
