package leetcode

import (
	"math"
	"sort"
)

func twoSumV2(target int, nums []int) [][]int {
	twoSum := make([][]int, 0)
	for l, r := 0, len(nums)-1; l < r; {
		if sum := nums[l] + nums[r]; sum == target {
			twoSum = append(twoSum, []int{nums[l], nums[r]})
			for l += 1; l < r && nums[l] == nums[l-1]; l++ {
			}
			for r -= 1; l < r && nums[r] == nums[r+1]; r-- {
			}
		} else if sum > target {
			r--
		} else {
			l++
		}
	}
	return twoSum
}

func threeSum(nums []int) [][]int {
	threeSum := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if twoSum := twoSumV2(0-nums[i], nums[i+1:]); len(twoSum) > 0 {
			for _, sum := range twoSum {
				threeSum = append(threeSum, []int{nums[i], sum[0], sum[1]})
			}
		}
	}
	return threeSum
}

var maximalDepth int

func maxDepthRe(root *TreeNode, curDepth int) {
	if root == nil {
		return
	}

	if curDepth++; curDepth > maximalDepth {
		maximalDepth = curDepth
	}

	maxDepthRe(root.Left, curDepth)
	maxDepthRe(root.Right, curDepth)
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maximalDepth = math.MinInt32
	maxDepthRe(root, 0)
	return maximalDepth
}

func letterCombinations(digits string) []string {

	if len(digits) == 0 {
		return []string{}
	}

	mapping := make(map[string][]string)
	mapping["2"] = []string{"a", "b", "c"}
	mapping["3"] = []string{"d", "e", "f"}
	mapping["4"] = []string{"g", "h", "i"}
	mapping["5"] = []string{"j", "k", "l"}
	mapping["6"] = []string{"m", "n", "o"}
	mapping["7"] = []string{"p", "q", "r", "s"}
	mapping["8"] = []string{"t", "u", "v"}
	mapping["9"] = []string{"w", "x", "y", "z"}

	res := []string{""}
	for i := 0; i < len(digits); i++ {
		letters := mapping[digits[i:i+1]]
		outerCombines := make([]string, 0)
		for _, letter := range letters {
			innerCombines := make([]string, 0)
			for _, re := range res {
				re += letter
				innerCombines = append(innerCombines, re)
			}
			outerCombines = append(outerCombines, innerCombines...)
		}
		res = outerCombines
	}
	return res
}
