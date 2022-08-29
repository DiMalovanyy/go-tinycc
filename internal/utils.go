package internal

import (
	"C"
	"runtime"
)
import "unsafe"

// Will set at compile time
var TccHeaders string

func LocateCHeaders() string {
	//TODO: Change this logic, not hardcode
	if runtime.GOOS == "linux" {
		return "/usr/include"
	}
	return ""
}

func CPointerToInterface(ptr *C.void) (*interface{}) {
	return (*interface{})(unsafe.Pointer(ptr))
}
