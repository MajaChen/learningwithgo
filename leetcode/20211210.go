package leetcode

import "math"

type MyQueue struct {
	stacks []*Stack
}

func QueueConstructor() MyQueue {

	return MyQueue{stacks: []*Stack{
		&Stack{elems: make([]interface{}, 0)},
		&Stack{elems: make([]interface{}, 0)},
	}}
}

func (this *MyQueue) Push(x int) {

	this.stacks[0].Push(x)
	return
}

func (this *MyQueue) transfer() {

	for !this.stacks[0].IsEmpty() {
		this.stacks[1].Push(this.stacks[0].Pop())
	}
}

func (this *MyQueue) Pop() int {

	if this.stacks[1].IsEmpty() {
		this.transfer()
	}
	return this.stacks[1].Pop().(int)
}

func (this *MyQueue) Peek() int {

	if this.stacks[1].IsEmpty() {
		this.transfer()
	}
	return this.stacks[1].GetTop().(int)
}

func (this *MyQueue) Empty() bool {

	return this.stacks[0].IsEmpty() && this.stacks[1].IsEmpty()
}

func isUgly(n int) bool {

	if n == 1 || n == 2 || n == 3 || n == 5 {
		return true
	}

	var flag bool
	for i := 2; i <= int(math.Ceil(math.Sqrt(float64(n)))); i++ {
		if n%i == 0 {
			flag = true
			j := n / i
			if !isUgly(i) || !isUgly(j) {
				return false
			}
		}
	}

	return flag
}

func isUgly2(n int) bool {

	if n == 0 {
		return false
	}

	var an int
	for an = n; an%2 == 0; an /= 2 {
	}
	for ; an%3 == 0; an /= 3 {
	}
	for ; an%5 == 0; an /= 5 {
	}

	return an == 1
}

// 从（i，j）出发，看是否可以突围
func solveRecur(board [][]byte, i, j int) {

	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || board[i][j] != 79 {
		return
	}

	board[i][j] = 65
	solveRecur(board, i-1, j)
	solveRecur(board, i+1, j)
	solveRecur(board, i, j-1)
	solveRecur(board, i, j+1)

	return
}

// 从四周向中间突破
func solve2(board [][]byte) {

	if len(board) <= 0 {
		return
	}

	n, m := len(board), len(board[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 || j == 0 || i == n-1 || j == m-1 {
				solveRecur(board, i, j)
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == 79 {
				board[i][j] = 88
			} else if board[i][j] == 65 {
				board[i][j] = 79
			}
		}
	}
	return
}
