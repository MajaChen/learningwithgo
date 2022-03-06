package leetcode

import (
	"sort"
	"strings"
)

func trailingZeroes(n int) int {
	//f(n) = n/5+n/5^2+n/5^3
	i, count := 5, 0
	for ; n/i > 0; i *= 5 {
		count += n / i
	}
	return count
}

var combinations2 [][]int

func combinationSumRe2(candidates, acombinations []int, target int) {
	if target == 0 {
		bcombinations := make([]int, len(acombinations))
		copy(bcombinations, acombinations)
		combinations2 = append(combinations2, bcombinations)
		return
	}
	if target < 0 {
		return
	}
	for i := 0; i < len(candidates); {
		initLength := len(acombinations)
		acombinations = append(acombinations, candidates[i])
		combinationSumRe2(candidates[i+1:], acombinations, target-candidates[i])
		acombinations = acombinations[:initLength]
		for i++; i < len(candidates) && candidates[i] == candidates[i-1]; i++ {
		}
	}
	return
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	combinations2 = make([][]int, 0)
	combinationSumRe2(candidates, make([]int, 0), target)
	return combinations
}

func swapElems(i, j, ai, aj int, matrix [][]int) {
	num := matrix[i][j]
	matrix[i][j] = matrix[ai][aj]
	matrix[ai][aj] = num
}

func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			ai, aj := n-j-1, n-i-1
			swapElems(i, j, ai, aj, matrix)
		}
	}

	for i := 0; i < n/2; i++ {
		for j := 0; j < n; j++ {
			ai, aj := n-i-1, j
			swapElems(i, j, ai, aj, matrix)
		}
	}
}

func transfer(str string) string {
	strs := make([]string, 0, len(str))
	for i := 0; i < len(str); i++ {
		strs = append(strs, str[i:i+1])
	}
	sort.Strings(strs)
	return strings.Join(strs, "-")
}

func groupAnagrams(strs []string) [][]string {

	keyWordSet, strSet := make([]string, 0), make([][]string, 0)
	for _, str := range strs {
		var hit bool
		keyWord := transfer(str)
		for i := 0; i < len(keyWordSet); i++ {
			if keyWord == keyWordSet[i] {
				strSet[i] = append(strSet[i], str)
				hit = true
				break
			}
		}

		if !hit {
			keyWordSet = append(keyWordSet, keyWord)
			strSet = append(strSet, []string{str})
		}
	}
	return strSet
}
