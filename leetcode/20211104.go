package leetcode

import "sort"

var prev *TreeNode
var newRoot *TreeNode

func BiNode(root *TreeNode) {
	if root == nil {
		return
	}

	if root.Left == nil && newRoot == nil {
		newRoot = root
	}

	BiNode(root.Left)
	root.Left = nil
	if prev != nil {
		prev.Right = root
	}
	prev = root
	BiNode(root.Right)
}

func twoSum(nums []int, target int) []int {
	mapping := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := mapping[target-nums[i]]; ok {
			return []int{mapping[target-nums[i]], i}
		} else {
			mapping[nums[i]] = i
		}
	}
	return []int{-1, -1}
}

func isValid(s string) bool {

	mapping := make(map[string]int)
	mapping["("] = 1
	mapping["["] = 2
	mapping["{"] = 3
	mapping[")"] = -1
	mapping["]"] = -2
	mapping["}"] = -3

	stack := Stack{}
	for i := 0; i < len(s); i++ {
		if mapping[s[i:i+1]] > 0 {
			stack.Push(mapping[s[i:i+1]])
		} else {
			if stack.IsEmpty() || stack.Pop().(int)+mapping[s[i:i+1]] != 0 {
				return false
			}
		}
	}

	return stack.IsEmpty()
}

func getKthMagicNumber2(k int) int {
	if k < 1 {
		return -1
	}
	nums := make([]int, 0, k)
	nums = append(nums, 1)
	k -= 1

	elems := []int{3, 5, 7}

	set := make(map[int]bool)
	index := 0
	for k > 0 {
		for _, elem := range elems {
			tail := nums[index]
			if _, ok := set[elem*tail]; !ok {
				nums = append(nums, elem*tail)
				sort.Ints(nums)
				set[elem*tail] = true
				if k--; k == 0 {
					return nums[len(nums)-1]
				}
			}
		}
		index++
	}
	return nums[len(nums)-1]
}
