package leetcode

//leetcode submit region begin(Prohibit modification and deletion)

var board [][]byte
var visitedMapping [][]byte
var word string

func dfsTraverse(i, j, pos int) bool {

	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) ||
		visitedMapping[i][j] == 1 {
		return false
	}

	if pos >= len(word) {
		return true
	}

	if byte(board[i][j]) != byte(word[pos]) {
		return false
	}

	visitedMapping[i][j] = 1
	flag := dfsTraverse(i+j, j, pos+1) ||
		dfsTraverse(i-1, j, pos+1) ||
		dfsTraverse(i, j+1, pos+1) ||
		dfsTraverse(i, j-1, pos+1)
	visitedMapping[i][j] = 0
	return flag
}

func exist(board [][]byte, word string) bool {

	m, n := len(board), len(board[0])
	board, word, visitedMapping = board, word, make([][]byte, 0, m)
	for i := 0; i < len(board); i++ {
		board = append(board, make([]byte, n))
	}

	remaing := m * n
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if remaing < len(word) {
				return false
			}
			if dfsTraverse(i, j, 0) {
				return true
			}
		}
	}

	return false
}

//leetcode submit region end(Prohibit modification and deletion)
