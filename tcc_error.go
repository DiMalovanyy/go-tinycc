package tinycc

import "errors"

var (
	ErrTccCouldNotCompileString = errors.New("tcc could not compile C code from given buffer")
	ErrTccCouldNotAddFile = errors.New("tcc could not add file to compilation context")
)
