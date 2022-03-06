package leetcode

import (
	"sort"
	"strconv"
	"strings"
)

func invertTreeRe(root *TreeNode) {
	if root == nil {
		return
	}

	invertTreeRe(root.Left)
	invertTreeRe(root.Right)
	node := root.Left
	root.Left = root.Right
	root.Right = node
}

func invertTree(root *TreeNode) *TreeNode {
	invertTreeRe(root)
	return root
}

// 从给定的startIndex开始，选取count个元素，考虑所有情况
func subsetsWithDupRe(nums []int, count, startIndex int) [][]int {
	if count == 0 {
		return [][]int{[]int{}}
	}

	res := make([][]int, 0)
	for i := startIndex; i+count <= len(nums); {
		subres := subsetsWithDupRe(nums, count-1, i+1)
		for j := 0; j < len(subres); j++ {
			re := append([]int{nums[i]}, subres[j]...)
			subres[j] = re
		}
		res = append(res, subres...)

		for i++; i+count <= len(nums) && nums[i] == nums[i-1]; i++ {
		}
	}
	return res
}

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	for span := 0; span <= len(nums); span++ {
		res = append(res, subsetsWithDupRe(nums, span, 0)...)
	}
	return res
}

func canDecode(s string, startIndex, span int) bool {

	if startIndex+span > len(s) || strings.HasPrefix(s[startIndex:startIndex+span], "0") {
		return false
	}

	if num, err := strconv.Atoi(s[startIndex : startIndex+span]); err != nil || num <= 0 || num > 26 {
		return false
	}

	return true
}

var decodeMapping []int

// 从startIndex出发，有多少种解释方式
func numDecodingsRe(s string, startIndex int) int {

	if startIndex >= len(s) {
		return 1
	}

	if decodeMapping[startIndex] >= 0 {
		return decodeMapping[startIndex]
	}

	count := 0
	if canDecode(s, startIndex, 1) {
		count += numDecodingsRe(s, startIndex+1)
	}
	if canDecode(s, startIndex, 2) {
		count += numDecodingsRe(s, startIndex+2)
	}

	decodeMapping[startIndex] = count
	return count
}

func numDecodings(s string) int {
	decodeMapping = make([]int, 0, len(s))
	for i := 0; i < len(s); i++ {
		decodeMapping = append(decodeMapping, -1)
	}
	return numDecodingsRe(s, 0)
}
