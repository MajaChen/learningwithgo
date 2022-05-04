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
func deleteNode(head *ListNode, val int) *ListNode {
	newHead := &ListNode{Next: head}
	p, q := newHead, newHead.Next
	for q != nil && q.Val != val {
		p = q
		q = q.Next
	}
	if q != nil {
		p.Next = q.Next
	}
	return newHead.Next
}

//leetcode submit region end(Prohibit modification and deletion)
