package tinycc

// Compile libtcc.a
//go:generate make

// #cgo LDFLAGS: -L${SRCDIR}/lib/tinycc -ltcc -ldl
// #cgo CFLAGS: -I${SRCDIR}/lib/tinycc
/*
#include <libtcc.h>
#include <stdint.h>
	static void* convert_int_to_addr(uint64_t addr) {
		return (void*)((uintptr_t)addr);
	}
*/
import (
	"C"
)
import (
	"unsafe"

	"github.com/DiMalovanyy/go-tinycc/internal"
)

// Context that describes TCC state
type TccContext struct {
	State *C.TCCState
}


// Creates new Tcc Compilation Context
func NewTccContext() (*TccContext, error) {
	//TODO: Link and include standard lib sources
	context := &TccContext {
		State: C.tcc_new(),
	}

	context.AddIncludePath(internal.LocateCHeaders())
	context.AddIncludePath(internal.TccHeaders)
	return context, nil
}

// Free a Tcc Compilation Context
func (c *TccContext) DeleteContext() {
	C.tcc_delete(c.State)
}

func (c *TccContext) SetErrorCallback() {
	//TODO:
}

/* -------------------- Preprocessor----------------------------------------*/

func (c *TccContext) AddIncludePath(includePath string) error {
	rc := C.tcc_add_include_path(
		c.State,
		(*C.char)(unsafe.Pointer(
			&([]byte(includePath))[0],
		)),
	)
	if rc == -1 {
		return ErrTccCouldNotAddIncludePath
	}
	return nil
}

// define preprocessor symbol 'sym'. value can be NULL, sym can be "sym=val"
func (c *TccContext) DefineSymbol(name string, address uint64) error {
	C.tcc_define_symbol(
		c.State,
		(*C.char)(unsafe.Pointer(
			&([]byte(name))[0],
		)),
		(*C.char)(
			unsafe.Pointer(C.convert_int_to_addr(
				(C.uint64_t)(address),
			)),
		),
	)
	// TODO: Handle C call failed (no Return code)
	return nil
}

// undefine preprocess symbol 'sym' 
func (c *TccContext) UndefineSymbol(name string) error {
	C.tcc_undefine_symbol(
		c.State,
		(*C.char)(unsafe.Pointer(
			&([]byte(name))[0],
		)),
	)
	// TODO: Handle C call failed (no Return code)
	return nil
}
/* -------------------------------------------------------------------------*/
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

// Compile C code located in buf
// NOTE: It only compile string (NOT Run)
func (c *TccContext) CompileString(buf []byte) error {
	rc := C.tcc_compile_string(c.State, (*C.char)(unsafe.Pointer(&buf[0])))
	if rc == -1 {
		return	ErrTccCouldNotCompileString
	}
	return nil
}

// Set option as from command line
func (c *TccContext) SetOption(option string) {
	//TODO:
}

// Set options as from command line
func (c *TccContext) SetOptions(options ...string) {
	for _, option := range options {
		c.SetOption(option)
	}
}


/* -------------------------------------------------------------------------*/
/* ---------------------- Linking ------------------------------------------*/

func (c *TccContext) SetOutputMode(mode OutputMode) error {
	if mode > 4 {
	   return ErrTccUndefinedOutputMode
	}
	rc := C.tcc_set_output_type(c.State, (C.int)(mode))
	if rc == -1 {
	   return ErrTccErrorOnOutputModeSet
	}
	return nil
}

func (c *TccContext) AddLibraryPath(libraryPath string) error {
	rc := C.tcc_add_library_path(
		c.State,
		(*C.char)(unsafe.Pointer(
			&([]byte(libraryPath))[0],
		)),
	)
	if rc == -1 {
		return ErrTccCouldNotAddLibraryPath
	}
	return nil
}

func (c *TccContext) AddLibrary(library string) error {
	rc := C.tcc_add_library(
		c.State,
		(*C.char)(unsafe.Pointer(
			&([]byte(library))[0],
		)),
	)
	if rc == -1 {
		return ErrTccCouldNotAddLibrary
	}
	return nil
}

// Add a symbol to compiled program (Dynamic add)
func (c *TccContext) AddSymbol(name string, address uint64) error {
	rc := C.tcc_add_symbol(
		c.State,
		(*C.char)(unsafe.Pointer(
			&([]byte(name))[0],
		)),
		unsafe.Pointer(C.convert_int_to_addr(
			(C.uint64_t)(address),
		)),
	)
	if rc == -1 {
		return ErrCouldNotAddSymbol
	}
	return nil
}

/* -------------------------------------------------------------------------*/
/* ---------------------- Executing ----------------------------------------*/

// Link and run main() and return it value
// NOTE: Do Not call Relocate() before Run()
func (c *TccContext) Run(args ...string) int { 
	return (int)(C.tcc_run(
		c.State,
		(C.int)(len(args)),
		(**C.char)(unsafe.Pointer(&args[0])),
	))
}

func (c *TccContext) Relocate() error {
	// TODO: !!!
	return nil
}


