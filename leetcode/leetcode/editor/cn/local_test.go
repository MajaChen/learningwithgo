package leetcode

import (
	"testing"
)

func TestXXX(t *testing.T) {
	t.Error(defangIPaddr("1.1.1.1"))
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
