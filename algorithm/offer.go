package algorithm

import (
	"math"
	"strings"
)

type CQueue struct {
	Input  *Stack
	Output *Stack
}

func Constructor() CQueue {
	return CQueue{Input: &Stack{elems: make([]interface{}, 0)},
		Output: &Stack{elems: make([]interface{}, 0)}}
}

func (this *CQueue) AppendTail(value int) {
	this.Input.Push(value)
}

func (this *CQueue) DeleteHead() int {
	if this.Output.IsEmpty() { //把input中的数据迁移到output
		if this.Input.IsEmpty() {
			return -1
		}
		for !this.Input.IsEmpty() {
			this.Output.Push(this.Input.Pop())
		}
		return this.DeleteHead()
	} else { //直接输出output栈顶元素
		return this.Output.Pop().(int)
	}
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */

func swap(i, j int, nums []int) {
	tmp := nums[i]
	nums[i] = nums[j]
	nums[j] = tmp
}

func findRepeatNumber(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for nums[i] != i {
			j := nums[i]
			if nums[j] == nums[i] {
				return nums[j]
			}
			swap(i, j, nums)
		}
	}
	return -1
}

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	rowBound, cloBound, i, j := len(matrix), len(matrix[0]), 0, 0
	for i != rowBound && j != cloBound {
		//行序遍历
		ai := j
		for ; ai < cloBound; ai++ {
			if matrix[i][ai] == target {
				return true
			} else if matrix[i][ai] > target {
				break
			}
		}
		if ai == cloBound {
			i++
		} else {
			cloBound = ai
		}
		//列序遍历
		aj := i
		for ; aj < rowBound; aj++ {
			if matrix[aj][j] == target {
				return true
			} else if matrix[aj][j] > target {
				break
			}
		}
		if aj == rowBound {
			j++
		} else {
			rowBound = aj
		}
	}
	return false
}

func firstIndex(arr []int, elem int) int {
	index := -1
	for i := 0; i < len(arr); i++ {
		if arr[i] == elem {
			index = i
			return index
		}
	}
	return index
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) != len(inorder) || len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}
	index := firstIndex(inorder, preorder[0])
	lInorder := inorder[:index]
	rInorder := inorder[index+1:]
	lPreorder := preorder[1 : 1+len(lInorder)]
	rPreorder := preorder[1+len(lInorder):]

	root.Left = buildTree(lPreorder, lInorder)
	root.Right = buildTree(rPreorder, rInorder)
	return root
}

func search(root *TreeNode, val int) []*TreeNode {
	nodes := make([]*TreeNode, 0)
	stack := Stack{elems: make([]interface{}, 0)}
	stack.Push(&Item{color: 0, Node: root})
	for !stack.IsEmpty() {
		if item := stack.Pop().(*Item); item.Node != nil {
			if colors[item.color] == "white" {
				stack.Push(&Item{color: 0, Node: item.Node.Right})
				stack.Push(&Item{color: 0, Node: item.Node.Left})
				stack.Push(&Item{color: 1, Node: item.Node})
			} else {
				if item.Node.Val == val {
					nodes = append(nodes, item.Node)
				}
			}
		}
	}
	return nodes
}

func isSubStructureRe(A *TreeNode, B *TreeNode) bool {
	lflag, rflag := true, true
	if B.Left != nil {
		if A.Left != nil && B.Left.Val == A.Left.Val {
			lflag = isSubStructureRe(A.Left, B.Left)
		} else {
			lflag = false
		}
	}
	if B.Right != nil {
		if A.Right != nil && B.Right.Val == A.Right.Val {
			rflag = isSubStructureRe(A.Right, B.Right)
		} else {
			rflag = false
		}
	}
	return lflag && rflag
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return false
	}
	//遍历A，寻找值与val对应的结点
	val := B.Val
	nodes := search(A, val)
	for _, node := range nodes {
		if isSubStructureRe(node, B) {
			return true
		}
	}
	return false
}

func exchangeNode(root *TreeNode) {
	tmp := root.Left
	root.Left = root.Right
	root.Right = tmp
}
func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if root.Left != nil {
		root.Left = mirrorTree(root.Left)
	}
	if root.Right != nil {
		root.Right = mirrorTree(root.Right)
	}
	exchangeNode(root)
	return root
}

func isArrSymmetric(arr []*TreeNode) bool {
	for i := 0; i < len(arr)/2; i++ {
		if arr[i] != nil && arr[len(arr)-1-i] != nil { //同时都是实心结点
			if arr[i].Val != arr[len(arr)-1-i].Val {
				return false
			}
		} else { // 有空心结点
			if arr[i] != arr[len(arr)-1-i] {
				return false
			}
		}
	}
	return true
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var q Queue = Queue{elems: make([]interface{}, 0)}
	q.Push(root)
	cur, next, count, tmpRes := 1, 0, 0, make([]*TreeNode, 0)
	for !q.IsEmpty() {
		for count < cur {
			node := q.Pop().(*TreeNode)
			tmpRes = append(tmpRes, node)
			if node != nil {
				q.Push(node.Left)
				next++
				q.Push(node.Right)
				next++
			}
			count++
		}
		if !isArrSymmetric(tmpRes) {
			return false
		}
		count = 0
		cur = next
		next = 0
		tmpRes = tmpRes[:0]
	}
	return true
}

func spiralOrder(matrix [][]int) []int {

	if len(matrix) == 0 {
		return []int{}
	}
	count, l := 0, len(matrix)*len(matrix[0])
	res := make([]int, 0, l)

	rowStart, colStart, rowBound, colBound := 0, 0, len(matrix)-1, len(matrix[0])-1
	for count < l {
		//向右
		for j := colStart; count < l && j <= colBound; j++ {
			res = append(res, matrix[rowStart][j])
			count++
		}
		rowStart++
		//向下
		for i := rowStart; count < l && i <= rowBound; i++ {
			res = append(res, matrix[i][colBound])
			count++
		}
		colBound--
		//向左
		for j := colBound; count < l && j >= colStart; j-- {
			res = append(res, matrix[rowBound][j])
			count++
		}
		rowBound--
		//向上
		for i := rowBound; count < l && i >= rowStart; i-- {
			res = append(res, matrix[i][colStart])
			count++
		}
		colStart++

	}
	return res[:len(matrix)*len(matrix[0])]
}

func exchange(nums []int) []int {
	l, r := -1, -1
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 { //偶数
			r = i
			if l < 0 {
				l = r
			}
		} else { //奇数
			if l >= 0 && r >= 0 { //左侧有偶数
				swap(i, l, nums)
				r = i
				l = l + 1
			}
		}
	}
	return nums
}

//是众数的即为+1，不是众数的记为-1
func majorityElement(nums []int) int {
	count, candidate := 0, 0
	for i := 0; i < len(nums); i++ {
		if count == 0 {
			candidate = nums[i]
		}
		if candidate == nums[i] {
			count++
		} else {
			count--
		}
	}
	return candidate
}

func replaceSpace(s string) string {
	sc := 0
	//首先统计空格的个数
	for i := 0; i < len(s); i++ {
		if s[i:i+1] == " " {
			sc++
		}
	}
	//然后计算新的字符串的长度
	nl := len(s) + sc*2

	//为新的字符串申请内存空间
	var sb strings.Builder
	sb.Grow(nl)
	for i := 0; i < len(s); i++ {
		if s[i:i+1] != " " {
			sb.WriteString(s[i : i+1])
		} else {
			sb.WriteString("%20")
		}
	}
	return sb.String()
}

func reversePrintRe(head *ListNode, re []int) []int {
	if head.Next == nil {
		re = append(re, head.Val)
		return re
	}

	re = reversePrintRe(head.Next, re)
	re = append(re, head.Val)
	return re
}

func reversePrint(head *ListNode) []int {
	re := make([]int, 0)
	if head == nil {
		return re
	}
	re = reversePrintRe(head, re)
	return re
}

// 对年龄排序
// 数据的特点是在某一个固定范围内反复重复
func sortAges(ages []int) []int {
	// 假设年龄的范围在1到120之间
	arr := make([]int, 121)
	for i := 0; i < len(ages); i++ {
		arr[ages[i]]++
	}
	count := 0
	for i := 0; i < len(arr); i++ {
		for j := 0; j < arr[i]; j++ {
			ages[count] = i
			count++
		}
	}
	return ages
}


func minArray(numbers []int) int {
	l, r := 0, len(numbers)-1
	for l <= r {
		mid := (l + r) / 2
		if numbers[mid] > numbers[r] { //左半部分元素不考虑
			l = mid + 1
		} else if numbers[mid] < numbers[r] { //右半部分元素不考虑
			r = mid
		} else {
			r = r - 1
		}
	}
	//取numbers[]
	return numbers[l]
}

func getLeastNumbers(arr []int, k int) []int {
	head, size := &Node{}, 0
	for i := 0; i < len(arr); i++ {
		p := head
		for p != nil {
			if p.Next == nil || arr[i] > p.Next.Val {
				p.Next = &Node{Val: arr[i], Next: p.Next}
				size++
				break
			}
			p = p.Next
		}
		if size > k {
			head.Next = head.Next.Next
		}
	}
	res, p := make([]int, 0, size), head.Next
	for p != nil {
		res = append(res, p.Val)
		p = p.Next
	}
	return res
}

//以递归的方式求1+2+3+...n
func addNums(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	return nums[0] + addNums(nums[1:])
}

func fib(n int) int {
	mapping := make(map[int]uint64)
	mapping[0], mapping[1] = 0, 1
	if n <= 1 {
		return int(mapping[n])
	}
	for i := 2; i <= n; i++ {
		mapping[i] = (mapping[i-1] + mapping[i-2]) % 1000000007
	}
	return int(mapping[n] % 1000000007)
}

func numWays(n int) int {
	mapping := make(map[int]uint64)
	mapping[0], mapping[1] = 1, 1
	if n <= 1 {
		return int(mapping[n])
	}
	for i := 2; i <= n; i++ {
		mapping[i] = (mapping[i-1] + mapping[i-2]) % 1000000007
	}
	return int(mapping[n] % 1000000007)
}

func hammingWeight(num uint32) int {

	i, count := uint32(1<<31), uint32(0)
	for j := 31; j >= 0; j-- {
		count += ((num << j) & i) >> 31
	}
	return int(count)
}

func hammingWeight2(num uint32) int {
	var count, flag uint32 = 0, 1
	for flag != 0 {
		if num&flag != 0 {
			count++
		}
		flag = flag << 1

	}
	return int(count)
}

func hammingWeight3(num uint32) int {
	count := 0
	for num != 0 {
		num = (num - 1) & num
		count++
	}
	return count
}

func printNumbers(n int) []int {
	max := math.Pow10(n) - 1
	res := make([]int, 0, int(max))
	for i := 1; i <= int(max); i++ {
		res = append(res, i)
	}
	return res
}

