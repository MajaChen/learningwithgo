package leetcode

import "testing"

func TestXXX(t *testing.T) {
	t.Error(buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}))
}

func TestMap(t *testing.T) {
	mapping := make(map[int]bool)
	mapping[0] = true
	t.Error(mapping[1])
}

func TestSlice(t *testing.T) {
	s := make([]int, 0, 2)
	s = append(s, 1)
	s = append(s, 2)
	t.Error(s[1:1])
}
