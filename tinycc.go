package tinycc

// Compile libtcc.a
//go:generate make

// #cgo LDFLAGS: -L${SRCDIR}/lib/tinycc -ltcc -ldl
// #cgo CFLAGS: -I${SRCDIR}/lib/tinycc
// #include <tcc.h>
import (
	"C"
)
import (
	"bytes"
	"unsafe"
)

// Context that describes TCC state
type TccContext struct {
	State *C.TCCState
}

// Creates new Tcc Compilation Context
func NewTccContext() (*TccContext, error) {
	return &TccContext {
		State: C.tcc_new(),
	}, nil
}

// Free a Tcc Compilation Context
func (c *TccContext) DeleteContext() {
	C.tcc_delete(c.State)
}


/* ---------------------- Compiling ----------------------------------------*/

// Adds a file (C file, dll, object, library, ld script).
func (c *TccContext) AddFile(filePath string) error {
	rc := C.tcc_add_file(
		c.State,
		(*C.char)(unsafe.Pointer(
			&([]byte(filePath))[0],
		)),
	)
	if rc == -1 {
		return ErrTccCouldNotAddFile
	}
	return nil
}

// Compile C code loacted in buf
func (c *TccContext) CompileString(buf *bytes.Buffer) error {
	bytes := buf.Bytes()
	rc := C.tcc_compile_string(c.State, (*C.char)(unsafe.Pointer(&bytes[0])))
	if rc == -1 {
		return	ErrTccCouldNotCompileString
	}
	return nil
}

/* -------------------------------------------------------------------------*/
/* ---------------------- Compiling ----------------------------------------*/






