package leetcode

import "strconv"

func moveZeroes(nums []int) {

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			for j := i - 1; j >= 0 && nums[j] == 0; j-- {
				swap(j+1, j, nums)
			}
		}
	}
}

func isPowerOfFour(n int) bool {

	if n == 0 {
		return false
	}
	for ; n%4 == 0; n /= 4 {
	}
	return n == 1
}

func preorderTraversal(root *TreeNode) []int {

	if root == nil {
		return []int{}
	}

	vals := make([]int, 0)
	colorMapping := map[int]string{0: "white", 1: "gray"}

	s := Stack{elems: make([]interface{}, 0)}
	s.Push(&Item{color: 0, Node: root})

	var item *Item

	for !s.IsEmpty() {
		if item = s.Pop().(*Item); item.Node != nil {
			if colorMapping[item.color] == "white" {
				s.Push(&Item{color: 0, Node: item.Node.Right})
				s.Push(&Item{color: 0, Node: item.Node.Left})
				s.Push(&Item{color: 1, Node: item.Node})
			} else {
				vals = append(vals, item.Node.Val)
			}
		}
	}
	return vals
}

func insert(head *ListNode, node *ListNode) {
	var p *ListNode
	for p = head; p.Next != nil && p.Next.Val <= node.Val; p = p.Next {
	}
	node.Next = p.Next
	p.Next = node
}

func insertionSortList(head *ListNode) *ListNode {
	vhead, nhead := &ListNode{Next: head}, &ListNode{}
	for p := vhead; p.Next != nil; {
		node := p.Next
		p.Next = p.Next.Next
		insert(nhead, node)
	}
	return nhead.Next
}

func evalRPN(tokens []string) int {

	numStack := &Stack{elems: make([]interface{}, 0)}
	operMapping := map[string]bool{"+": true, "-": true, "*": true, "/": true}

	for _, token := range tokens {
		if _, ok := operMapping[token]; ok {
			nums1, nums2 := numStack.Pop().(int), numStack.Pop().(int)
			if token == "+" {
				numStack.Push(nums1 + nums2)
			} else if token == "-" {
				numStack.Push(nums2 - nums1)
			} else if token == "*" {
				numStack.Push(nums1 * nums2)
			} else {
				numStack.Push(nums2 / nums1)
			}
		} else {
			num, _ := strconv.Atoi(token)
			numStack.Push(num)
		}
	}

	return numStack.Pop().(int)
}
