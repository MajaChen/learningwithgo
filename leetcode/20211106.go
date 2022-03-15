package leetcode

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

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	p, q := l1, l2
	newHead := &ListNode{}
	np := newHead
	flag := 0
	for p != nil && q != nil {
		var sum int
		if sum = flag + p.Val + q.Val; sum >= 10 {
			sum -= 10
			flag = 1
		} else {
			flag = 0
		}
		np.Next = &ListNode{Val: sum}
		np = np.Next
		p = p.Next
		q = q.Next
	}

	var tail *ListNode
	if flag == 1 {
		one := &ListNode{Val: 1}
		if p != nil {
			tail = addTwoNumbers2(one, p)
		} else if q != nil {
			tail = addTwoNumbers2(one, q)
		} else {
			tail = one
		}
	} else {
		if p != nil {
			tail = p
		} else if q != nil {
			tail = q
		}
	}
	np.Next = tail
	return newHead.Next
}

func removeDuplicates(nums []int) int {
	j := -1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			if j < 0 {
				j = i
			}
		} else {
			if i != j && j > 0 {
				nums[j] = nums[i]
				j++
			}
		}
	}

	if j < 0 {
		return len(nums)
	}
	return j
}
