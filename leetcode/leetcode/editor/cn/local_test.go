package leetcode

import (
	"math"
	"testing"
)

func TestXXX(t *testing.T) {
	t.Error(nextGreaterElement(2147483486))
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
	t.Error(2147483486 > math.MaxInt32)
	t.Error(2147483648 > math.MaxInt32)
	var i int64
	i = 2147483648
	t.Error(i)
}
