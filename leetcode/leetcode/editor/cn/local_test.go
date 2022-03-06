package leetcode

import "testing"

func TestXXX(t *testing.T) {
	t.Error(minimumEffortPath([][]int{
		{1, 2, 2},
		{3, 8, 2},
		{5, 3, 5},
	}))
}

func TestMap(t *testing.T) {
	mapping := make(map[int]bool)
	mapping[0] = true
	t.Error(mapping[1])
}
