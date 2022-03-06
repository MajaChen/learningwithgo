package leetcode

import (
	"sort"
	"testing"
)

func TestDecodeString(t *testing.T) {

}

func TestFor(t *testing.T) {
	var i, j int
	j = 1
	// 执行流程：
	// 1.i = 0;2.检验j>0;3.j -= 1;4.i++;5.检验j>0
	for i = 0; j > 0; i++ {
		j -= 1
	}
	t.Error(i) //1
}

func TestFor2(t *testing.T) {
	var i int
	for i = 0; i < 1; i++ {
	}
	t.Error(i) //1
}

func TestBinarySearch(t *testing.T) {

	nums := []int{1, 3, 7, 22}
	i := sort.Search(len(nums), func(i int) bool {
		return nums[i] >= 8
	})
	t.Error(i)
}
