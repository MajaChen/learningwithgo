package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type Stack struct {
	elems []interface{}
}

func (s *Stack) IsEmpty() bool {
	return len(s.elems) == 0
}

func (s *Stack) len() int {
	return len(s.elems)
}

func (s *Stack) Pop() interface{} {
	elem := s.elems[len(s.elems)-1]
	s.elems = s.elems[:len(s.elems)-1]
	return elem
}

func (s *Stack) Push(elem interface{}) {
	s.elems = append(s.elems, elem)
}

func (s *Stack) GetTop() interface{} {
	return s.elems[len(s.elems)-1]
}

func nextLargerNodes(head *ListNode) []int {

	mapping, stack, nodeCounts, count := make(map[int]int), &Stack{elems: make([]interface{}, 0)}, make(map[int]int), 1
	for p := head; p != nil; p = p.Next {
		mapping[count] = p.Val
		for !stack.IsEmpty() && mapping[stack.GetTop().(int)] < p.Val {
			nodeCounts[stack.Pop().(int)] = p.Val
		}
		stack.Push(count)
		count++
	}

	res := make([]int, count-1)
	for index, count := range nodeCounts {
		res[index-1] = count
	}

	return res
}

//leetcode submit region end(Prohibit modification and deletion)
