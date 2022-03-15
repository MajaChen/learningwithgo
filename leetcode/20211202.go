package leetcode

func containsNearbyDuplicate(nums []int, k int) bool {

	for mapping, i := make(map[int]int), 0; i < len(nums); i++ {
		if _, ok := mapping[nums[i]]; ok {
			if i-mapping[nums[i]] <= k {
				return true
			} else {
				mapping[nums[i]] = i
			}
		} else {
			mapping[nums[i]] = i
		}
	}
	return false
}

func partition2(head *ListNode, x int) *ListNode {
	newHead, newHeadA := &ListNode{Next: head}, &ListNode{}
	pa, p := newHeadA, newHead
	for p != nil && p.Next != nil {
		if p.Next.Val < x {
			// 将p.Next指向结点摘除
			// p指针不前移
			q := p.Next
			p.Next = p.Next.Next
			q.Next = nil
			// 将摘除结点插入新结点
			pa.Next = q
			pa = pa.Next
		} else {
			// p指针前移
			p = p.Next
		}
	}

	// 将 newHeadA、newHead、newHeadB连接起来
	pa.Next = newHead.Next
	return newHeadA.Next
}
