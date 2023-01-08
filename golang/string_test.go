package golang

import "testing"

func TestNilString(t *testing.T) {
	var s string
	if s != "" {
		t.Error("s is not blank")
	} else {
		t.Error("s is blank")
	}
}
