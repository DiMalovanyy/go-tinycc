package tinycc

import "testing"

func TestTccStateCreateDelete(t *testing.T) {
	tccContext, err := NewTccContext()
	if err != nil {
		t.Errorf("Could not create tccState: %v", err)
	}
	defer tccContext.DeleteContext()
}
