package leetcode

import (
	"testing"
)

func TestXXX(t *testing.T) {
	t.Error(sequenceReconstruction([]int{1, 2, 3}, [][]int{
		{1, 2},
		{1, 3},
		{2, 3},
	}))

}

func TestSlice(t *testing.T) {
	s := make([]int, 0, 2)
	s = append(s, 1)
	s = append(s, 2)
	t.Error(s[1:1])
}

func TestMap(t *testing.T) {
	mapping := make(map[int]int)
	if _, ok := mapping[0]; ok {
		t.Error("equal")
	}
}

func TestMax(t *testing.T) {
	m := make(map[int]int)
	m[2] += 1
	m[2] += 1
	t.Error(m)
}
