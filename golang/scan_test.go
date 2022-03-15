package golang

import "testing"

func TestScan(t *testing.T) {
	if err := Scan(); err != nil {
		t.Errorf("err is %v", err)
	}
}
