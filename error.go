package tinycc

import "errors"

var (
	ErrTccCouldNotAddIncludePath = errors.New("tcc could not add include path to compulation context")
	ErrTccCouldNotCompileString = errors.New("tcc could not compile C code from given buffer")
	ErrTccCouldNotAddFile = errors.New("tcc could not add file to compilation context")
	ErrTccUndefinedOutputMode = errors.New("provided undefined output mode as tcc linking option")
	ErrTccErrorOnOutputModeSet = errors.New("could not set output mode as tcc linking option")
	ErrTccCouldNotAddLibraryPath = errors.New("tcc could not add library path to linking option")
	ErrTccCouldNotAddLibrary = errors.New("tcc could not add path to linking option")
)
