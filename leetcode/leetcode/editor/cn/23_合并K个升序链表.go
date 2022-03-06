package leetcode

//leetcode submit region begin(Prohibit modification and deletion)
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

func mergeKListsRecursion(lists []*ListNode, l, r int) *ListNode {

	if l >= r {
		return lists[l]
	}

	mid := (l + r) >> 1
	l1, l2 := mergeKListsRecursion(lists, l, mid), mergeKListsRecursion(lists, mid+1, r)
	return mergeTwoLists(l1, l2)
}

func mergeKLists(lists []*ListNode) *ListNode {

	if len(lists) == 0 {
		return nil
	}
	return mergeKListsRecursion(lists, 0, len(lists)-1)
}

//leetcode submit region end(Prohibit modification and deletion)
