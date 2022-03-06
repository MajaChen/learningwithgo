package leetcode

import "sort"

func majorityElement(nums []int) int {
	var elem int
	for i, counter := 0, 0; i < len(nums); i++ {
		if counter == 0 {
			elem = nums[i]
			counter++
			continue
		}
		if nums[i] == elem {
			counter++
		} else {
			counter--
		}
	}
	return elem
}

var combinations [][]int

func combinationSumRe(candidates, acombinations []int, target int) {
	if target == 0 {
		bcombinations := make([]int, len(acombinations))
		copy(bcombinations, acombinations)
		combinations = append(combinations, bcombinations)
		return
	}
	if target < 0 {
		return
	}
	for i := 0; i < len(candidates); i++ {
		initLength := len(acombinations)
		for atarget := target; atarget >= candidates[i]; {
			atarget -= candidates[i]
			acombinations = append(acombinations, candidates[i])
			combinationSumRe(candidates[i+1:], acombinations, atarget)
		}
		acombinations = acombinations[:initLength]
	}
	return
}

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	combinations = make([][]int, 0)
	combinationSumRe(candidates, make([]int, 0), target)
	return combinations
}
