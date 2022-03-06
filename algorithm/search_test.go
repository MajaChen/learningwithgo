package algorithm

import "testing"

func TestK2IJ(t *testing.T) {
	k := 5
	row, col := k2IJ(k)
	t.Errorf("row and col are %v,%v\n", row, col)
}
