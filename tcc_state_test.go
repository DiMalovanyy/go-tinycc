package tinycc

// #cgo LDFLAGS: -L${SRCDIR}/lib/tinycc -ltcc
// #cgo CFLAGS: -I${SRCDIR}/lib/tinycc
import "testing"

func TestTccStateCreateDelete(t *testing.T) {
	tccState, err := NewTccState()
	if err != nil {
		t.Errorf("Could not create tccState: %v", err)
	}
	defer tccState.DeleteState()
}
