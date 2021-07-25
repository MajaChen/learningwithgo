package algorithm

func add(a, b int) int {
	return a + b
}

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	nhead := &ListNode{}
	p := nhead
	p1, p2 := l1, l2
	for p1 != nil && p2 != nil {
		if p1.Val <= p2.Val {
			p.Next = &ListNode{Val: p1.Val, Next: nil}
			p1 = p1.Next
		} else {
			p.Next = &ListNode{Val: p2.Val, Next: nil}
			p2 = p2.Next
		}
		p = p.Next
	}
	if p1 != nil {
		p.Next = p1
	}
	if p2 != nil {
		p.Next = p2
	}
	return nhead.Next
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	pre, cur := head, head.Next
	for cur != nil {
		if pre.Val == cur.Val { //删除cur指向元素,前移cur指针
			pre.Next = cur.Next
			cur = cur.Next
			if cur == nil {
				return nil
			}
		} else { // pre、cur指针同时前移
			pre = cur
			cur = cur.Next
		}
	}
	return head
}

func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	vhead := &ListNode{Val: 0, Next: nil}
	vhead.Next = head //增加一个虚拟结点
	pre, cur, next := vhead, head, head.Next
	for next != nil {
		if cur.Val == next.Val {
			for next.Next != nil && next.Val == next.Next.Val {
				next = next.Next
			}
			pre.Next = next.Next //删除重复元素
			if next != nil {
				cur = next.Next
				next = cur.Next
			}
		} else {
			pre = cur
			cur = next
			next = next.Next
		}
	}
	return vhead.Next
}

func index(head *ListNode, x int) *ListNode {
	count := 0
	p := head
	for p != nil {
		count++
		if count == x {
			return p
		}
		p = p.Next
	}
	return nil
}

func vindex(head *ListNode, x int) *ListNode {
	if x == 0 {
		return head
	}
	count := 0
	p := head.Next
	for p != nil {
		count++
		if count == x {
			return p
		}
		p = p.Next
	}
	return nil
}
func partition(head *ListNode, x int) *ListNode {
	xNode := index(head, x)
	if xNode == nil {
		return nil
	}
	pivot := xNode.Val
	p := head
	lhead := &ListNode{}
	rhead := &ListNode{}
	lp, rp := lhead, rhead
	for p != nil {
		if p.Val < pivot {
			lp.Next = &ListNode{Val: p.Val, Next: nil}
			lp = lp.Next
		} else {
			rp.Next = &ListNode{Val: p.Val, Next: nil}
			rp = rp.Next
		}
		p = p.Next
	}
	lp.Next = rhead.Next
	return lhead.Next
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	//加一个虚拟的头
	vhead := &ListNode{}
	vhead.Next = head

	lnode := vindex(vhead, left)
	rnode := vindex(vhead, right)
	//left结点的前驱结点
	lpnode := vindex(vhead, left-1)
	//right结点的后继结点
	ranode := rnode.Next
	if lnode == nil || rnode == nil {
		return head
	}
	if lnode == rnode || lnode.Next == rnode {
		return head
	}
	var pre *ListNode = nil
	cur := lnode
	var aNext *ListNode
	for cur != ranode {
		aNext = cur.Next
		cur.Next = pre
		pre = cur //pre不需要管
		cur = aNext
	}

	lpnode.Next = rnode
	lnode.Next = ranode
	return vhead.Next
}

func sortListRe(head *ListNode) *ListNode {
	//统计长度
	p, count := head, 0
	for p != nil {
		count++
		p = p.Next
	}
	var len int
	if len = count; len <= 1 {
		return head
	}

	// 分治
	lhead := head
	var rhead *ListNode = nil
	p, count = head, 1
	for count != len/2 {
		count++
		p = p.Next
	}
	rhead = p.Next
	p.Next = nil
	lhead = sortListRe(lhead)
	rhead = sortListRe(rhead)

	// 合并
	return mergeTwoLists(lhead, rhead)
}

func sortList(head *ListNode) *ListNode {
	return sortListRe(head)
}

func reorderList(head *ListNode) {
	p, count, arr := head, 0, make([]*ListNode, 0)
	for p != nil {
		count++
		arr = append(arr, p)
		p = p.Next
	}
	len := count

	nhead := &ListNode{}
	p = nhead
	for i := 0; i <= len/2; i++ {
		p.Next = arr[i]
		p.Next.Next = arr[len-i-1]
		p = p.Next.Next
	}

	if len%2 == 0 { //长度是偶数
		p.Next = nil
	} else { //长度是奇数
		p.Next = arr[len/2]
		p.Next.Next = nil
	}
	head = nhead
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	p1, p2 := head.Next, head.Next.Next
	for p2 != nil {
		if p1 == p2 {
			return true
		}
		if p2.Next == nil {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next.Next
	}
	return false
}

//Definition for a Node.
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {

	p, count, nmapping, omapping, nhead := head, 0, make(map[int]*Node), make(map[*Node]int), &Node{}
	omapping[nil] = -1
	nmapping[-1] = nil 
	np := nhead
	for p != nil {
		np.Next = &Node{Val: p.Val, Next: nil, Random: p.Random}
		nmapping[count] = np.Next
		omapping[p] = count
		np = np.Next
		p = p.Next
		count++
	}

	np = nhead.Next
	for np != nil {
		np.Random = nmapping[omapping[np.Random]]
		np = np.Next
	}
	return nhead.Next
}

func detectCycle(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return nil
	}
	p1, p2 := head, head
	//首先让快指针和慢指针首次相遇
	for p2.Next.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
		if p1 == p2 { //然后让慢指针走往前走a步
			p := head
			for p != p1 {
				p = p.Next
				p1 = p1.Next
			}
			return p
		}
	}
	return nil
}


