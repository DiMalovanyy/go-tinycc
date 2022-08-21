package tinycc

// Compile libtcc.a
//go:generate make

// #cgo LDFLAGS: -L${SRCDIR}/lib/tinycc -ltcc -ldl
// #cgo CFLAGS: -I${SRCDIR}/lib/tinycc
// #include <tcc.h>
import (
	"C"
)

type TccState struct {
	State *C.TCCState
}

func NewTccState() (*TccState, error) {
	return &TccState {
		State: C.tcc_new(),
	}, nil
}

func (s *TccState) DeleteState() {
	C.tcc_delete(s.State)
}
