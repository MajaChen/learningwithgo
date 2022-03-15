package algorithm

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Errorf("testing result is %v", add(1, 2))
}

func TestReverseBetween(t *testing.T) {
	head := &ListNode{Val: 1, Next: nil}
	node2 := &ListNode{Val: 2, Next: nil}
	node3 := &ListNode{Val: 3, Next: nil}
	node4 := &ListNode{Val: 4, Next: nil}
	node5 := &ListNode{Val: 5, Next: nil}
	head.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	head = reverseBetween(head, 1, 4)
	printList(head)
}

func TestPartition(t *testing.T) {
	head := &ListNode{Val: 1, Next: nil}
	node4 := &ListNode{Val: 4, Next: nil}
	node3 := &ListNode{Val: 3, Next: nil}
	node2 := &ListNode{Val: 2, Next: nil}
	node5 := &ListNode{Val: 5, Next: nil}
	node2_2 := &ListNode{Val: 2, Next: nil}
	head.Next = node4
	node4.Next = node3
	node3.Next = node2
	node2.Next = node5
	node5.Next = node2_2
	head = partition(head, 3)
	printList(head)
}

func printList(head *ListNode) {
	p := head
	for p != nil {
		fmt.Printf("%v ", p.Val)
		p = p.Next
	}
}

func TestDeleteDuplicates(t *testing.T) {
	head := &ListNode{Val: 1, Next: nil}
	node1 := &ListNode{Val: 1, Next: nil}
	node2 := &ListNode{Val: 2, Next: nil}
	head.Next = node1
	node1.Next = node2
	head = deleteDuplicates(head)
	printList(head)
}

func TestDeleteDuplicates2(t *testing.T) {
	head := &ListNode{Val: 1, Next: nil}
	node1 := &ListNode{Val: 1, Next: nil}
	node2 := &ListNode{Val: 2, Next: nil}
	node3 := &ListNode{Val: 3, Next: nil}
	node3_3 := &ListNode{Val: 3, Next: nil}
	node4 := &ListNode{Val: 4, Next: nil}
	node4_4 := &ListNode{Val: 4, Next: nil}
	node5 := &ListNode{Val: 5, Next: nil}
	head.Next = node1
	node1.Next = node2
	node2.Next = node3
	node3.Next = node3_3
	node3_3.Next = node4
	node4.Next = node4_4
	node4_4.Next = node5

	head = deleteDuplicates2(head)
	printList(head)
}

func TestMergeTwoLists(t *testing.T) {
	head1 := &ListNode{Val: 1, Next: nil}
	node1_2 := &ListNode{Val: 2, Next: nil}
	node1_4 := &ListNode{Val: 4, Next: nil}
	head1.Next = node1_2
	node1_2.Next = node1_4

	head2 := &ListNode{Val: 1, Next: nil}
	node2_3 := &ListNode{Val: 3, Next: nil}
	node2_4 := &ListNode{Val: 4, Next: nil}
	head2.Next = node2_3
	node2_3.Next = node2_4

	head := mergeTwoLists(head1, head2)
	printList(head)
}

func TestSortList(t *testing.T) {
	head := &ListNode{Val: 4, Next: nil}
	node2 := &ListNode{Val: 2, Next: nil}
	node1 := &ListNode{Val: 1, Next: nil}
	node3 := &ListNode{Val: 3, Next: nil}
	head.Next = node2
	node2.Next = node1
	node1.Next = node3
	head = sortList(head)
	printList(head)
}

func TestHasCycle(t *testing.T) {
	head := &ListNode{Val: 4, Next: nil}
	node2 := &ListNode{Val: 2, Next: nil}
	node1 := &ListNode{Val: 1, Next: nil}
	node3 := &ListNode{Val: 3, Next: nil}
	head.Next = node2
	node2.Next = node1
	node1.Next = node3

	t.Errorf("has cycly :%v", hasCycle(head))
}

func TestDetectCycle(t *testing.T) {
	head := &ListNode{Val: 4, Next: nil}
	node2 := &ListNode{Val: 2, Next: nil}
	node1 := &ListNode{Val: 1, Next: nil}
	node3 := &ListNode{Val: 3, Next: nil}
	head.Next = node2
	node2.Next = node1
	node1.Next = node3
	//	node3.Next = node2

	t.Errorf("entrance point is %v", detectCycle(head))
}

func TestReorderList(t *testing.T) {
	head := &ListNode{Val: 4, Next: nil}
	node2 := &ListNode{Val: 2, Next: nil}
	node1 := &ListNode{Val: 1, Next: nil}
	node3 := &ListNode{Val: 3, Next: nil}
	node5 := &ListNode{Val: 5, Next: nil}

	head.Next = node2
	node2.Next = node1
	node1.Next = node3
	node3.Next = node5

	reorderList(head)
	printList(head)
}

func TestCopyRandomList(t *testing.T) {

	one := Node{Val: 1, Next: nil, Random: nil}
	ten := Node{Val: 10, Next: &one, Random: nil}
	eleven := Node{Val: 11, Next: &ten, Random: nil}
	thirteen := Node{Val: 13, Next: &eleven, Random: nil}
	seven := Node{Val: 7, Next: &thirteen, Random: nil}

	thirteen.Random = &seven
	eleven.Random = &one
	ten.Random = &eleven
	one.Random = &seven

	list := copyRandomList(&seven)
	t.Errorf("list is %v", list)
}

func TestMergeKLists(t *testing.T) {
	headA := &ListNode{Val: 1}
	three := &ListNode{Val: 3}
	headA.Next = three

	headB := &ListNode{Val: 2}
	five := &ListNode{Val: 5}
	headB.Next = five

	list := []*ListNode{headA, headB}
	newList := mergeKLists(list)
	printList(newList)
}
