package algorithm

import (
	"math"
	"strconv"
)

//  Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Item struct {
	color int //0 白色，1 灰色
	Node  *TreeNode
}

var colors map[int]string

func init() {
	colors = make(map[int]string)
	colors[0] = "white"
	colors[1] = "gray"
}

type Stack struct {
	elems []interface{}
}

func (s *Stack) IsEmpty() bool {
	return len(s.elems) == 0
}

func (s *Stack) Len() int {
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

func PreTraverse(root *TreeNode) []int {
	res := make([]int, 0)
	stack := Stack{elems: make([]interface{}, 0)}
	stack.Push(&Item{color: 0, Node: root})
	for !stack.IsEmpty() {
		if item := stack.Pop().(*Item); item.Node != nil {
			if colors[item.color] == "white" {
				stack.Push(&Item{color: 0, Node: item.Node.Right})
				stack.Push(&Item{color: 0, Node: item.Node.Left})
				stack.Push(&Item{color: 1, Node: item.Node})
			} else {
				res = append(res, item.Node.Val)
			}
		}
	}
	return res

}

func PostTraverse(root *TreeNode) []int {
	res := make([]int, 0)
	stack := Stack{elems: make([]interface{}, 0)}
	stack.Push(&Item{color: 0, Node: root})
	for !stack.IsEmpty() {
		if item := stack.Pop().(*Item); item.Node != nil {
			if colors[item.color] == "white" {
				stack.Push(&Item{color: 1, Node: item.Node})
				stack.Push(&Item{color: 0, Node: item.Node.Right})
				stack.Push(&Item{color: 0, Node: item.Node.Left})
			} else {
				res = append(res, item.Node.Val)
			}
		}
	}
	return res
}

type Queue struct {
	elems []interface{}
}

func (q *Queue) Size() int {
	return len(q.elems)
}

func (q *Queue) Pop() interface{} {
	elem := q.elems[0]
	q.elems = q.elems[1:]
	return elem
}

func (q *Queue) Push(elem interface{}) {
	q.elems = append(q.elems, elem)
}

func (q Queue) IsEmpty() bool {
	return len(q.elems) == 0
}

func layerTraverse(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	var q Queue = Queue{elems: make([]interface{}, 0)}
	q.Push(root)
	for !q.IsEmpty() {
		node := q.Pop().(*TreeNode)
		res = append(res, node.Val)
		if node.Left != nil {
			q.Push(node.Left)
		}
		if node.Right != nil {
			q.Push(node.Right)
		}
	}
	return res
}

// 得到每一层的结点，二维数组的第一维标识节点层次
func getNodesPerLayer(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	var q Queue = Queue{elems: make([]interface{}, 0)}
	q.Push(root)
	count, tmpRes := 0, make([]int, 0)
	for !q.IsEmpty() {
		size := q.Size()
		for count < size {
			node := q.Pop().(*TreeNode)
			tmpRes = append(tmpRes, node.Val)
			if node.Left != nil {
				q.Push(node.Left)
			}
			if node.Right != nil {
				q.Push(node.Right)
			}
			count++
		}
		count = 0
		res = append(res, tmpRes)
		//tmpRes = tmpRes[:0] 禁止内存复用，否则会出错
		tmpRes = make([]int, 0, q.Size())
	}
	return res
}

func bstFromPreorder(preorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}

	i := 1
	for ; i < len(preorder) && preorder[i] <= root.Val; i++ {
	}
	lPreorder := preorder[1:i]
	rPreorder := preorder[i:]
	root.Left = bstFromPreorder(lPreorder)
	root.Right = bstFromPreorder(rPreorder)
	return root
}

func distanceNode(root *TreeNode, distance int, visited map[*TreeNode]bool) []int {

	res := make([]int, 0)
	if root == nil {
		return res
	}
	var q Queue = Queue{elems: make([]interface{}, 0)}
	q.Push(root)
	count, tmpRes, d := 0, make([]int, 0), 0
	for !q.IsEmpty() {
		size := q.Size()
		for count < size {
			node := q.Pop().(*TreeNode)
			tmpRes = append(tmpRes, node.Val)
			if node.Left != nil && !visited[node.Left] {
				q.Push(node.Left)
			}
			if node.Right != nil && !visited[node.Right] {
				q.Push(node.Right)
			}
			count++
		}
		count = 0
		if d++; d == distance+1 {
			return tmpRes
		}
		tmpRes = make([]int, 0, q.Size())
	}
	return res
}

func parent(root, target *TreeNode) *TreeNode {
	if root == nil || target == nil {
		return nil
	}
	if root.Left == target || root.Right == target {
		return root
	}

	if p := parent(root.Left, target); p != nil {
		return p
	}
	if p := parent(root.Right, target); p != nil {
		return p
	}
	return nil
}

func distanceK(root *TreeNode, target *TreeNode, k int) []int {

	visited := make(map[*TreeNode]bool)
	visited[target] = true

	//遍历子结点
	res := distanceNode(target, k, visited)
	//遍历target的上层节点
	nk, p := k, target
	for nk >= 0 {
		if p = parent(root, p); p == nil {
			return res
		}
		nk--

		node := distanceNode(p, nk, visited)
		visited[p] = true
		res = append(res, node...)
	}

	return res
}

func InTraverseRe(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	res = append(res, InTraverseRe(root.Left)...)
	res = append(res, root.Val)
	res = append(res, InTraverseRe(root.Right)...)
	return res
}

func isValidBST(root *TreeNode) bool {
	res := InTraverseRe(root)
	for i := 1; i < len(res); i++ {
		if res[i] < res[i-1] {
			return false
		}
	}
	return true
}

func nodeExist(nodeIndex int, root *TreeNode) bool {
	paths, p := strconv.FormatInt(int64(nodeIndex), 2), root
	for i := 1; i < len(paths) && p != nil; i++ {
		if paths[i] == 48 {
			p = p.Left
		} else {
			p = p.Right
		}
	}
	return p != nil
}

func countNodes(root *TreeNode) int {

	if root == nil {
		return 0
	}

	// 首先求出结点层数h
	p, h := root, -1
	for p != nil {
		p = p.Left
		h++
	}

	// 然后计算出结点的上下限
	lowBound := int(math.Pow(2, float64(h)))
	upperBound := int(math.Pow(2, float64(h+1))) - 1

	//然后寻找第一个不存在的结点
	l, r := lowBound, upperBound
	for l < r { //循环结束条件是l==r
		mid := (l + r) / 2
		if nodeExist(mid, root) { //去掉mid连同左边部分
			l = mid + 1
		} else { //去掉右边部分
			r = mid
		}
	}
	if nodeExist(l, root) {
		return l
	} else {
		return l - 1
	}
}

/*
func getNodesPerLayer(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	var q Queue = Queue{elems: make([]type_func{}, 0)}
	q.Push(root)
	count, tmpRes := 0, make([]int, 0)
	for !q.IsEmpty() {
		size := q.Size()
		for count < size {
			node := q.Pop().(*TreeNode)
			tmpRes = append(tmpRes, node.Val)
			if node.Left != nil {
				q.Push(node.Left)
			}
			if node.Right != nil {
				q.Push(node.Right)
			}
			count++
		}
		count = 0
		res = append(res, tmpRes)
		//tmpRes = tmpRes[:0] 禁止内存复用，否则会出错
		tmpRes = make([]int, 0, q.Size())
	}
	return res
}

*/

// 给定二叉树的根结点，求其深度
func depth(root *TreeNode) int {

	if root == nil {
		return 0
	}

	var q Queue = Queue{elems: make([]interface{}, 0)}
	q.Push(root)
	count, depth := 0, 0
	for !q.IsEmpty() {
		size, hasValNode := q.Size(), false
		for count < size {
			node := q.Pop().(*TreeNode)
			if node.Left != nil {
				hasValNode = true
				q.Push(node.Left)
			}
			if node.Right != nil {
				hasValNode = true
				q.Push(node.Right)
			}
			count++
		}
		depth++
		if !hasValNode {
			return depth
		}
		count = 0
	}
	return depth //永远都不会走到
}

//给定二叉树的根结点，序列化该二叉树
func serilize(root *TreeNode) []string {
	res := make([]string, 0)
	if root == nil {
		return res
	}

	var q Queue = Queue{elems: make([]interface{}, 0)}
	q.Push(root)
	count := 0
	var np *TreeNode = nil
	for !q.IsEmpty() {
		size, hasValNode, tmpRes := q.Size(), false, make([]string, 0)
		for count < size {
			node := q.Pop().(*TreeNode)
			if node != nil {
				hasValNode = true
				q.Push(node.Left)
				q.Push(node.Right)
				tmpRes = append(tmpRes, strconv.Itoa(node.Val))
			} else {
				q.Push(np) //此处不能用nil代替，否则类型断言会失败
				q.Push(np)
				tmpRes = append(tmpRes, "nil")
			}
			count++
		}

		if !hasValNode {
			return res
		}
		res = append(res, tmpRes...)
		count = 0
	}
	return res
}

var emptyTreeNodePointer *TreeNode

func buildTreeNode(str string) *TreeNode {
	if str == "nil" {
		return emptyTreeNodePointer
	}
	i, _ := strconv.Atoi(str)
	return &TreeNode{Val: i}
}

//给定二叉树序列化得到的字符串数组，反序列化出该二叉树
func deserilize(strs []string) *TreeNode {
	root, nodes := 0, make([]*TreeNode, 0, len(strs)+1)
	for root < len(strs) {
		node := buildTreeNode(strs[root])
		pos := root + 1
		nodes[pos] = node
		if (pos)/2 > 0 {
			p := nodes[(pos)/2]
			if (pos % 2) == 0 {
				p.Left = node
			} else {
				p.Right = node
			}
		}
	}
	return nodes[1]
}

var maxPathSumNum int = math.MinInt32

//局部最大值/向上传递的局部最大值
func maxPathSumRe(root *TreeNode) int {
	if root == nil {
		return 0
	}

	lval := maxPathSumRe(root.Left)
	rval := maxPathSumRe(root.Right)
	val, localmax := root.Val, root.Val
	if lval >= 0 {
		val += lval
	}
	if rval >= 0 {
		val += rval
	}

	if val > maxPathSumNum {
		maxPathSumNum = val
	}

	if lval > 0 && lval >= rval {
		localmax += lval
	}

	if rval > 0 && rval >= lval {
		localmax += rval
	}

	return localmax
}

func maxPathSum(root *TreeNode) int {
	maxPathSumRe(root)
	return maxPathSumNum
}

func combinePaths(pre int, lres [][]int, rres [][]int) [][]int {
	res := make([][]int, 0, len(lres)+len(rres))
	for _, tmpRes := range lres {
		tmpRes = append([]int{pre}, tmpRes...)
		res = append(res, tmpRes)
	}
	for _, tmpRes := range rres {
		tmpRes = append([]int{pre}, tmpRes...)
		res = append(res, tmpRes)
	}
	return res
}

func pathSum(root *TreeNode, targetSum int) [][]int {

	// 返回的res长度为0则表明没有路径
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	if root.Left == nil && root.Right == nil && root.Val == targetSum {
		res = append(res, []int{root.Val})
		return res
	}

	targetSum -= root.Val
	lres := pathSum(root.Left, targetSum)
	rres := pathSum(root.Right, targetSum)
	res = combinePaths(root.Val, lres, rres)

	return res
}

// 双递归法，第一层递归
func pathSum_v2(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	res := pathSum_double(root, sum)
	res += pathSum_v2(root.Left, sum)
	res += pathSum_v2(root.Right, sum)
	return res
}

// 双递归法，第二层递归
func pathSum_double(root *TreeNode, sum int) int {

	if root == nil {
		return 0
	}

	if sum == root.Val {
		return 1
	}

	sum -= root.Val
	return pathSum_double(root.Left, sum) + pathSum_double(root.Right, sum)
}

func findNodeTilt(root *TreeNode) int {

	lsum, rsum := 0, 0
	if root.Left != nil {
		lsum = root.Left.Val
	}
	if root.Right != nil {
		rsum = root.Right.Val
	}
	return int(math.Abs(float64(lsum - rsum)))
}

func findTilt(root *TreeNode) int {
	if root == nil {
		return 0
	}
	tilt := findNodeTilt(root)
	tilt += findTilt(root.Left)
	tilt += findTilt(root.Right)
	return tilt
}

func goodNodesRe(root *TreeNode) []int {
	gN := make([]int, 0)
	if root == nil {
		return gN
	}

	gN = append(gN, goodNodesRe(root.Left)...)
	gN = append(gN, goodNodesRe(root.Right)...)

	nGN := make([]int, 0, len(gN))
	for _, node := range gN {
		if node >= root.Val {
			nGN = append(nGN, node)
		}
	}
	nGN = append(nGN, root.Val)
	return nGN
}

func goodNodes(root *TreeNode) int {
	return len(goodNodesRe(root))
}

func nums2Int(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += int(math.Pow(float64(2), float64(len(nums)-i-1))) * nums[i]
	}
	return sum
}

var sum int = 0

func sumRootToLeafRe(root *TreeNode, nums []int) {
	if root == nil {
		return
	}

	nums = append(nums, root.Val)
	if root.Left == nil && root.Right == nil {
		sum += nums2Int(nums)
	}

	sumRootToLeafRe(root.Left, nums)
	sumRootToLeafRe(root.Right, nums)
}

func sumRootToLeaf(root *TreeNode) int {
	nums := make([]int, 0)
	sumRootToLeafRe(root, nums)
	return sum
}

// 存在1则返回true，否则返回false
func removeOneRe(root *TreeNode) bool {
	if root == nil {
		return false
	}

	if root.Val == 1 {
		return true
	}

	lflag := removeOneRe(root.Left)
	if !lflag {
		root.Left = nil
	}
	rflag := removeOneRe(root.Right)
	if !rflag {
		root.Right = nil
	}
	return lflag || rflag || root.Val == 1
}

func removeOne(root *TreeNode) *TreeNode {
	//建一个虚拟头结点,统一处理逻辑
	vRoot := &TreeNode{Val: 0, Left: root}
	removeOneRe(vRoot)
	return vRoot.Left
}

// 如果root应该被删除就返回true
func removeTargetRe(root *TreeNode, target int) bool {
	if root == nil {
		return true
	}

	var lflag, rflag bool
	if lflag = removeTargetRe(root.Left, target); lflag {
		root.Left = nil
	}
	if rflag = removeTargetRe(root.Right, target); rflag {
		root.Right = nil
	}

	if lflag && rflag && root.Val == target {
		return true
	}
	return false
}

func removeTarget(root *TreeNode, target int) *TreeNode {
	// 构建虚拟结点统一处理逻辑
	vRoot := &TreeNode{Val: target, Left: root}
	removeOneRe(vRoot)
	return vRoot.Left
}
