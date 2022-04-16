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
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	if l1 == nil && l2 != nil {
		return l2
	}

	p := &ListNode{Next: l1}
	head := p

	q := l2
	for q != nil {
		// 寻找q在l1中的插入位置
		for ; p.Next != nil && p.Next.Val <= q.Val; p = p.Next {
		}

		// 插入q到l1
		aq := q.Next
		q.Next = p.Next
		p.Next = q
		q = aq
	}

	return head.Next
}

//leetcode submit region end(Prohibit modification and deletion)
