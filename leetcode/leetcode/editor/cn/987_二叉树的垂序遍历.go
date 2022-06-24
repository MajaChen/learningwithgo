package leetcode

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//leetcode submit region begin(Prohibit modification and deletion)

type LeetcodeTreeNode struct {
	val int
	row int
	col int
}

var colNodesMapping map[int][]*LeetcodeTreeNode
var maximalCol int

type NodesSlice []*LeetcodeTreeNode

func (s NodesSlice) Len() int {
	return len(s)
}

func (s NodesSlice) Less(i, j int) bool {

	if s[i].row == s[j].row {
		return s[i].val < s[j].val
	} else {
		return s[i].row < s[j].row
	}
}

func (s NodesSlice) Swap(i, j int) {

	t := s[i]
	s[i] = s[j]
	s[j] = t
}

type ColNodes struct {
	vals []int
	col  int
}

type ColNodesSlice []*ColNodes

func (s ColNodesSlice) Len() int {
	return len(s)
}

func (s ColNodesSlice) Less(i, j int) bool {
	return s[i].col < s[j].col
}

func (s ColNodesSlice) Swap(i, j int) {

	t := s[i]
	s[i] = s[j]
	s[j] = t
}

func inOrderTraverse(root *TreeNode, row, col int) {

	if root == nil {
		return
	}

	node := &LeetcodeTreeNode{val: root.Val, row: row, col: col}
	var nodes []*LeetcodeTreeNode
	if _, ok := colNodesMapping[col]; ok {
		nodes = colNodesMapping[col]
	} else {
		nodes = make([]*LeetcodeTreeNode, 0)
	}
	nodes = append(nodes, node)
	colNodesMapping[col] = nodes

	inOrderTraverse(root.Left, row+1, col-1)
	inOrderTraverse(root.Right, row+1, col+1)
}

func verticalTraversal(root *TreeNode) [][]int {

	colNodesMapping = make(map[int][]*LeetcodeTreeNode)
	colNodesSlice := make([]*ColNodes, 0)
	inOrderTraverse(root, 0, 0)

	for col, nodes := range colNodesMapping {
		sort.Sort(NodesSlice(nodes))
		vals := make([]int, 0, len(nodes))
		for _, node := range nodes {
			vals = append(vals, node.val)
		}
		colNodesSlice = append(colNodesSlice, &ColNodes{vals: vals, col: col})
	}

	sort.Sort(ColNodesSlice(colNodesSlice))

	res := make([][]int, 0, len(colNodesSlice))
	for _, colNodes := range colNodesSlice {
		res = append(res, colNodes.vals)
	}

	return res
}

//leetcode submit region end(Prohibit modification and deletion)
