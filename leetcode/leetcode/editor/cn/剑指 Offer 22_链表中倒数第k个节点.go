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
func getKthFromEnd(head *ListNode, k int) *ListNode {
	if k <= 0 || head == nil {
		return head
	}

	var q *ListNode
	for p, counter := head, 1; p != nil; {
		if q != nil {
			q = q.Next
		}
		if counter == k {
			q = head
		}
		p = p.Next
		counter++
	}

	return q
}

//leetcode submit region end(Prohibit modification and deletion)
