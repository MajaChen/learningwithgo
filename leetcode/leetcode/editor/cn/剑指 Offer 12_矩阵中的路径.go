package leetcode

//leetcode submit region begin(Prohibit modification and deletion)

var traverseBoard [][]byte
var visitedMapping [][]byte
var traverseWord string

func dfsTraverse(i, j, pos int) bool {

	if i < 0 || i >= len(traverseBoard) || j < 0 || j >= len(traverseBoard[0]) ||
		visitedMapping[i][j] == 1 ||
		string(traverseBoard[i][j]) != traverseWord[pos:pos+1] {
		return false
	}

	if pos == len(traverseWord)-1 && byte(traverseBoard[i][j]) == byte(traverseWord[pos]) {
		return true
	}

	visitedMapping[i][j] = 1
	flag := dfsTraverse(i+1, j, pos+1) ||
		dfsTraverse(i-1, j, pos+1) ||
		dfsTraverse(i, j+1, pos+1) ||
		dfsTraverse(i, j-1, pos+1)
	visitedMapping[i][j] = 0
	return flag
}

func exist(board [][]byte, word string) bool {

	m, n := len(board), len(board[0])
	if m*n < len(word) {
		return false
	}

	traverseBoard, traverseWord, visitedMapping = board, word, make([][]byte, 0, m)
	for i := 0; i < len(board); i++ {
		visitedMapping = append(visitedMapping, make([]byte, n))
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if dfsTraverse(i, j, 0) {
				return true
			}
		}
	}

	return false
}

//leetcode submit region end(Prohibit modification and deletion)
