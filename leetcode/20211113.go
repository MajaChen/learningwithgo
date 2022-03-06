package leetcode

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	mid := (0 + len(nums)) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = sortedArrayToBST(nums[:mid])
	root.Right = sortedArrayToBST(nums[mid+1:])
	return root
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var p, q *ListNode
	var count int
	newHead := &ListNode{Next: head}
	for p, count = head, 0; p != nil; p = p.Next {
		if q != nil {
			q = q.Next
		}
		if count++; count == n+1 {
			q = head
		}
		// q指向的结点正好是倒数第n个结点的前一个结点
		if p.Next == nil && count >= 3 {
			q.Next = q.Next.Next
		}
	}

	if q == nil && n == count {
		newHead.Next = newHead.Next.Next
	}

	return newHead.Next
}
