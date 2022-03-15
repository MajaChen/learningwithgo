package algorithm

import (
	"testing"
)

func TestCQueue(t *testing.T) {
	cq := ConstructCQueue()
	cq.AppendTail(1)
	cq.AppendTail(2)
	cq.AppendTail(3)

	t.Error(cq.DeleteHead())
	cq.AppendTail(5)
	t.Error(cq.DeleteHead())
	t.Error(cq.DeleteHead())
	cq.AppendTail(6)
	cq.AppendTail(7)
	t.Error(cq.DeleteHead())
	t.Error(cq.DeleteHead())
	t.Error(cq.DeleteHead())
	t.Error(cq.DeleteHead())
	t.Error(cq.DeleteHead())
}

func TestFindRepeatNumber(t *testing.T) {
	arr := []int{2, 3, 1, 0, 5, 3}
	i := findRepeatNumber(arr)
	t.Error(i)
}

func TestFindNumberIn2DArray(t *testing.T) {
	doubleArr := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	i := findNumberIn2DArray(doubleArr, -1)
	t.Error(i)
}

func TestBuildTree(t *testing.T) {
	preOrder := []int{3, 9, 20, 15, 7}
	inOrder := []int{9, 3, 15, 20, 7}
	root := buildTree(preOrder, inOrder)
	res := layerTraverse(root)
	t.Errorf("res is %v", res)
}

func TestBuildTree2(t *testing.T) {
	preOrder := []int{3, 9, 20, 15, 7}
	inOrder := []int{7, 15, 20, 9, 3}
	root := buildTree(preOrder, inOrder)
	res := layerTraverse(root)
	t.Errorf("res is %v", res)
}

func TestIsSubStructure(t *testing.T) {
	one := TreeNode{Val: 1}
	two := TreeNode{Val: 2}
	four := TreeNode{Val: 4, Left: &one, Right: &two}
	five := TreeNode{Val: 5}
	three := TreeNode{Val: 3, Left: &four, Right: &five}

	one_2 := TreeNode{Val: 1}
	four_2 := TreeNode{Val: 4, Left: &one_2}

	is := isSubStructure(&three, &four_2)
	t.Error(is)
}

func TestSpiralOrder(t *testing.T) {
	matrix := [][]int{
		{1},
		{2},
		{3},
	}
	res := spiralOrder(matrix)
	t.Errorf("res is %v", res)
}

func TestExchange(t *testing.T) {
	nums := []int{}
	nums = exchange(nums)
	t.Errorf("nums is %v", nums)
}

func TestMajorityElement(t *testing.T) {
	nums := []int{1, 2, 3, 2, 2, 2, 5, 4, 2}
	i := majorityElement(nums)
	t.Error(i)
}

func TestReplaceSpace(t *testing.T) {
	s := "We are happy"
	s = replaceSpace(s)
	t.Error(s)
}

func TestReversePrint(t *testing.T) {
	head := &ListNode{Val: 1, Next: nil}
	node2 := &ListNode{Val: 2, Next: nil}
	node3 := &ListNode{Val: 3, Next: nil}
	node4 := &ListNode{Val: 4, Next: nil}
	node5 := &ListNode{Val: 5, Next: nil}
	head.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	re := reversePrint(head)
	t.Errorf("re is %v", re)
}

func TestSortAges(t *testing.T) {
	ages := []int{22, 23, 22, 42, 22}
	ages = sortAges(ages)
	t.Errorf("ages are %v", ages)
}

func TestIsSymmetric(t *testing.T) {
	left := TreeNode{Val: 2}
	right := TreeNode{Val: 2}
	root := TreeNode{Val: 1, Left: &left, Right: &right}
	t.Error(isSymmetric(&root))
}

func TestMinArray(t *testing.T) {
	numbers := []int{3, 4, 5, 1, 2}
	t.Error(minArray(numbers))
}

func TestGetLeastNumbers(t *testing.T) {
	arr, k := []int{4, 5, 1, 6, 2, 7, 3, 8}, 1
	t.Errorf("res is %v", getLeastNumbers(arr, k))
}

func TestGetLeastNumbers2(t *testing.T) {
	arr, k := []int{4}, 1
	t.Errorf("res is %v", getLeastNumbers(arr, k))
}

func TestAddNums(t *testing.T) {
	nums := []int{1, 2, 5, 7}
	res := addNums(nums)
	t.Error(res)
}

func TestFib(t *testing.T) {
	t.Error(fib(5))
}

func TestHammingWeight(t *testing.T) {
	n := 3
	hammingWeight2(uint32(n))
}
