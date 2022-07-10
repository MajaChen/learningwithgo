package leetcode

//leetcode submit region begin(Prohibit modification and deletion)

type LeetcodeListNode struct {
	Val  int
	Next *LeetcodeListNode
}

func insertList(head *LeetcodeListNode, insertVal int) *LeetcodeListNode {
	if head == nil {
		p := &LeetcodeListNode{Val: insertVal}
		p.Next = p
		return p
	}

	p, maxNode := head, head
	for {
		if p.Next.Val >= insertVal && p.Val <= insertVal {
			q := &LeetcodeListNode{Val: insertVal, Next: p.Next}
			p.Next = q
			return head
		}

		if p = p.Next; p == head {
			break
		} else if p.Val >= maxNode.Val {
			maxNode = p
		}
	}

	q := &LeetcodeListNode{Val: insertVal, Next: maxNode.Next}
	maxNode.Next = q
	return head
}

//leetcode submit region end(Prohibit modification and deletion)
