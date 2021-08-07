package golang

import "testing"

func TestSliceInFunction(t *testing.T) {
	nums := []int{1}
	SliceInFunction(nums)
	t.Errorf("new nums is %v", nums)
}
