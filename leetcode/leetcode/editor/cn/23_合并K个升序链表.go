package leetcode

import . "learning/common"

//leetcode submit region begin(Prohibit modification and deletion)

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
