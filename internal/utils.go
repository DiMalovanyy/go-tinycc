package internal

import "runtime"

// Will set at compile time
var TccHeaders string

func LocateCHeaders() string {
	//TODO: Change this logic, not hardcode
	if runtime.GOOS == "linux" {
		return "/usr/include"
	}
	return ""
}
