package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func TestXXX(t *testing.T) {
	t.Error(eraseOverlapIntervals([][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{1, 3},
	}))
}

func TestSlice(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1[0:3]

	s3 := []int{1, 2, 3, 4, 5}
	s4 := s3[0:3]
	fmt.Println(reflect.DeepEqual(s2, s4))
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
