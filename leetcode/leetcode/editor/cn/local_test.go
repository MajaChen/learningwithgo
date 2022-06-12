package leetcode

import (
	"testing"
)

func TestXXX(t *testing.T) {
	s := Constructor([][]int{
		[]int{-2, -2, 1, 1},
		[]int{2, 2, 4, 6},
	})
	t.Error(s.Pick())
	t.Error(s.Pick())
	t.Error(s.Pick())
	t.Error(s.Pick())
	t.Error(s.Pick())
	t.Error(s.Pick())
	t.Error(s.Pick())
	t.Error(s.Pick())
	t.Error(s.Pick())
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
