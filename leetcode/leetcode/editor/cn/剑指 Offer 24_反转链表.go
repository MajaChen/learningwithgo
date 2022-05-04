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
func reverseList(head *ListNode) *ListNode {
	var emptyNode *ListNode
	if head == emptyNode {
		return head
	}
	pp, p, q := emptyNode, head, head.Next
	for p != emptyNode {
		q = p.Next
		p.Next = pp
		pp = p
		p = q
	}
	return pp
}

//leetcode submit region end(Prohibit modification and deletion)
