package leetcode

import "testing"

func TestXXX(t *testing.T) {

	board := [][]byte{
		[]byte{'A', 'B', 'C', 'E'},
		[]byte{'S', 'F', 'C', 'S'},
		[]byte{'A', 'D', 'E', 'E'},
	}
	exist(board, "")

	findRepeatNumber([]int{})
}

func TestMap(t *testing.T) {
	mapping := make(map[int]bool)
	mapping[0] = true
	t.Error(mapping[1])
}
