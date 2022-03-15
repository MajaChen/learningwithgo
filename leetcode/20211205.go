package leetcode

func reverseList(head *ListNode) (*ListNode, *ListNode) {
	var pre, p, q *ListNode
	for p, q = head, head.Next; p != nil; {
		q = p.Next
		p.Next = pre
		pre = p
		p = q
	}
	return pre, head
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	newHead := &ListNode{Next: head}
	var reverseHead, reverseTail *ListNode
	for count, flag, p := 0, 0, newHead; flag != 2; p = p.Next {
		if count+1 == left {
			reverseHead = p
			flag++
		}
		if count == right {
			reverseTail = p
			flag++
		}
		count++
	}
	aAhead, aNext := reverseHead, reverseTail.Next
	reverseTail.Next = nil

	reverseHead, reverseTail = reverseList(reverseHead.Next)
	aAhead.Next, reverseTail.Next = reverseHead, aNext
	return newHead.Next
}

var colors map[int]string

func init() {
	colors = make(map[int]string)
	colors[0] = "white"
	colors[1] = "gray"
}

// 三色标记法，实现二叉树的迭代遍历
func InTraverse(root *TreeNode) []int {
	res := make([]int, 0)
	stack := Stack{elems: make([]interface{}, 0)}
	stack.Push(&Item{color: 0, Node: root})
	for !stack.IsEmpty() {
		if item := stack.Pop().(*Item); item.Node != nil {
			if colors[item.color] == "white" {
				stack.Push(&Item{color: 0, Node: item.Node.Right})
				stack.Push(&Item{color: 1, Node: item.Node})
				stack.Push(&Item{color: 0, Node: item.Node.Left})
			} else {
				res = append(res, item.Node.Val)
			}
		}
	}
	return res
}
