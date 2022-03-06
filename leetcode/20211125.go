package leetcode

import (
	"math"
	"strconv"
)

func factorial(n int) int {
	multi := 1
	for i := 1; i <= n; i++ {
		multi *= i
	}
	return multi
}

func nums2String(nums []int) (str string) {
	for i := len(nums) - 1; i >= 0; i-- {
		str += strconv.Itoa(nums[i])
	}
	return
}

func removeElem(nums []int, index int) []int {
	for i := index; i < len(nums)-1; i++ {
		nums[i] = nums[i+1]
	}
	return nums[:len(nums)-1]
}

func getPermutation(n int, k int) string {
	nums, permutationStr := make([]int, 0, n), ""
	for i := 0; i < n; i++ {
		nums = append(nums, i+1)
	}

	for k >= 0 {
		factor := factorial(len(nums) - 1)
		if k < factor {
			permutationStr += strconv.Itoa(nums[0])
			nums = removeElem(nums, 0)
		} else if k > factor {
			index := int(math.Ceil(float64(k)/float64(factor)) - 1)
			permutationStr += strconv.Itoa(nums[index])
			nums = removeElem(nums, index)
			k -= index * factor
		} else {
			permutationStr += strconv.Itoa(nums[0])
			permutationStr += nums2String(nums[1:])
			return permutationStr
		}
	}
	return permutationStr
}

func (head *ListNode) length() int {
	count := 0
	for p := head; p != nil; p = p.Next {
		count++
	}
	return count
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	n := head.length()
	if i := k % n; i == 0 {
		return head
	}
	var p, q *ListNode
	var count int
	for p = head; p.Next != nil; {
		p = p.Next
		if q != nil {
			q = q.Next
		}
		if count++; count == k {
			q = head
		}
	}

	newHead := q.Next
	q.Next = nil
	p.Next = head
	return newHead
}

var pathsMapping [][]int

func uniquePathsRe(m, n, i, j int) int {
	if i == m && j == n {
		return 1
	}

	if pathsMapping[i][j] >= 0 {
		return pathsMapping[i][j]
	}

	//尝试向右
	var leftPathsSum, rightPathsSum int
	if i+1 <= m {
		leftPathsSum = uniquePathsRe(m, n, i+1, j)
	}

	// 尝试向下
	if j+1 <= n {
		rightPathsSum = uniquePathsRe(m, n, i, j+1)
	}
	pathsMapping[i][j] = leftPathsSum + rightPathsSum
	return leftPathsSum + rightPathsSum
}

func uniquePaths(m int, n int) int {
	pathsMapping = make([][]int, 0)
	for i := 0; i < m; i++ {
		nums := make([]int, 0, n)
		for j := 0; j < n; j++ {
			nums[j] = -1
		}
		pathsMapping = append(pathsMapping, nums)
	}
	return uniquePathsRe(m-1, n-1, 0, 0)
}
