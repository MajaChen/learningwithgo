package algorithm

import (
	"testing"
	"time"
)

func TestArrangeNums(t *testing.T) {
	strs := arrangeNums(3)
	t.Errorf("res is %v", strs)
}

func TestArrangeAlphabet(t *testing.T) {
	strs := arrangeAlphabet(2)
	t.Errorf("strs is %v", strs)
}

func TestRemoveZ(t *testing.T) {
	s := "000123"
	s = removeZ(s)
	t.Errorf("s is %v", s)
}

func TestMap(t *testing.T) {
	mapping := make(map[string]bool)
	mapping["A"] = true

	nmapping := mapping
	nmapping["B"] = true
}

func TestSlice_v2(t *testing.T) {
	slices := make([][]string, 0)
	slices = append(slices, []string{"a"})
	for index, s := range slices {
		s = append(s, "b")
		slices[index] = s
	}

	t.Errorf("slices are %v",slices)
}

func TestReward(t *testing.T){
	candidates := []string{"A","B","C","D","E","F","G","H"}
	res := arrange(candidates,3,make(map[string]bool))
	t.Errorf("len of res is %v",len(res))
	t.Errorf("res is %v",res)
}

func TestCombineStr(t *testing.T){
	strs := []string{"A","B","C","D","E","F"}
	res := combine(strs,3)
	t.Errorf("len of res is %v",len(res))
	t.Errorf("res is %v",res)
}

func TestArrangeNumsNoRe(t *testing.T){
	start := time.Now()
	arrangeNumsNoRe(6)
	t.Errorf("len of time span is %v",time.Since(start))
}

func TestArrangeNums_v2(t *testing.T) {
	start := time.Now()
	arrangeNums(4)
	t.Errorf("len of time span is %v",time.Since(start))
}



