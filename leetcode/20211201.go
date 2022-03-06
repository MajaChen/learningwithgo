package leetcode

// 从给定的startIndex开始，选取count个元素，考虑所有情况
func subsetsRe(nums []int, count, startIndex int) [][]int {
	if count == 0 {
		return [][]int{[]int{}}
	}

	res := make([][]int, 0)
	for i := startIndex; i+count <= len(nums); i++ {
		subres := subsetsRe(nums, count-1, i+1)
		for j := 0; j < len(subres); j++ {
			re := append([]int{nums[i]}, subres[j]...)
			subres[j] = re
		}
		res = append(res, subres...)
	}
	return res
}

func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	res = append(res, subsetsRe(nums, 0, 0)...)
	for span := 1; span <= len(nums); span++ {
		res = append(res, subsetsRe(nums, span, 0)...)
	}
	return res
}

var visitedMapping [][]bool

func existRe(board [][]byte, word string, startIndex, i, j int) bool {

	char := byte([]rune(word[startIndex : startIndex+1])[0])
	if (i < 0 || i >= len(board) || j < 0 || j >= len(board[0])) ||
		visitedMapping[i][j] ||
		char != board[i][j] {
		return false
	}

	if startIndex == len(word)-1 && char == board[i][j] {
		return true
	}

	visitedMapping[i][j] = true
	flag := existRe(board, word, startIndex+1, i-1, j) || existRe(board, word, startIndex+1, i+1, j) ||
		existRe(board, word, startIndex+1, i, j-1) || existRe(board, word, startIndex+1, i, j+1)
	visitedMapping[i][j] = false

	return flag
}

func exist(board [][]byte, word string) bool {

	visitedMapping = make([][]bool, 0, len(board))
	for i := 0; i < len(board); i++ {
		visitedMapping = append(visitedMapping, make([]bool, len(board[0])))
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if existRe(board, word, 0, i, j) {
				return true
			}
		}
	}

	return false
}

func removeDuplicates2(nums []int) int {
	index := -1
	for count, currentElem, i := 1, nums[0], 1; i < len(nums); i++ {
		if nums[i] == currentElem {
			if count++; count > 2 {
				if index < 0 {
					index = i
				}
			} else {
				if index >= 0 {
					swap(index, i, nums)
					index++
				}
			}
		} else {
			count, currentElem = 1, nums[i]
			if index >= 0 {
				swap(index, i, nums)
				index++
			}

		}
	}
	if index < 0 {
		index = len(nums)
	}
	return index
}
